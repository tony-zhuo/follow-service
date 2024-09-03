package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	pkgKafka "github.com/tony-zhuo/follow-service/pkg/kafka"
	"github.com/tony-zhuo/follow-service/service/model"
	"sync"
	"time"
)

var (
	followUc     *FollowUc
	followUcOnce sync.Once
)

type FollowUc struct {
	dbRepo      model.FollowDBRepoInterface
	cacheRepo   model.FollowCacheRepoInterface
	kafkaWriter *kafka.Writer
}

func NewFollowUc(dbRepo model.FollowDBRepoInterface, cacheRepo model.FollowCacheRepoInterface, conf *pkgKafka.Config) *FollowUc {
	topic := "follow"
	prepareCfg := pkgKafka.NewConfig(conf.KafkaURL, topic, "", conf.SASLEnable, conf.UserName, conf.Password)

	followUcOnce.Do(func() {
		followUc = &FollowUc{
			dbRepo:      dbRepo,
			cacheRepo:   cacheRepo,
			kafkaWriter: pkgKafka.NewKafkaWriter(prepareCfg),
		}
	})
	return followUc
}

// Follow 關注
func (uc *FollowUc) Follow(ctx context.Context, req *model.FollowRequest) error {
	now := time.Now()
	// cache
	if err := uc.cacheRepo.Follow(ctx, &model.Follow{
		FollowerId: req.FollowerId,
		FolloweeId: req.FolloweeId,
		CreatedAt:  now,
	}); err != nil {
		// logging
		return err
	}

	// kafkax
	key := fmt.Sprintf("%s_%s_%s_%d", model.Action_Follow, req.FollowerId, req.FolloweeId, now.UnixNano())
	data, _ := json.Marshal(req)
	msg := pkgKafka.NewMsg("", []byte(key), data)
	if err := uc.kafkaWriter.WriteMessages(ctx, msg); err != nil {
		// logging
		_ = uc.cacheRepo.UnFollow(ctx, &model.Follow{
			FollowerId: req.FollowerId,
			FolloweeId: req.FolloweeId,
			CreatedAt:  now,
		})
		return err
	}

	return nil
}

// UnFollow 取消關注
func (uc *FollowUc) UnFollow(ctx context.Context, req *model.FollowRequest) error {
	now := time.Now()
	// cache
	if err := uc.cacheRepo.UnFollow(ctx, &model.Follow{
		FollowerId: req.FollowerId,
		FolloweeId: req.FolloweeId,
		CreatedAt:  now,
	}); err != nil {
		// logging
		return err
	}

	// kafka
	key := fmt.Sprintf("%s_%s_%s_%d", model.Action_UnFollow, req.FollowerId, req.FolloweeId, now.UnixNano())
	data, _ := json.Marshal(req)
	msg := pkgKafka.NewMsg("", []byte(key), data)
	if err := uc.kafkaWriter.WriteMessages(ctx, msg); err != nil {
		// logging
		_ = uc.cacheRepo.Follow(ctx, &model.Follow{
			FollowerId: req.FollowerId,
			FolloweeId: req.FolloweeId,
			CreatedAt:  now,
		})
		return err
	}

	return nil
}

// Followers 粉絲清單（分頁）
func (uc *FollowUc) Followers(ctx context.Context, req *model.SearchFollowerCond) ([]*model.Follow, error) {
	cacheData, err := uc.cacheRepo.Followers(ctx, req)
	if err != nil {
		// logging
		return nil, err
	}

	if cacheData != nil {
		return cacheData, nil
	}

	data, err := uc.dbRepo.Followers(ctx, req)
	if err != nil {
		// logging
		return nil, err
	}

	key := fmt.Sprintf("%s:%s", model.CacheKey_Followers, req.UserId)
	if err := uc.cacheRepo.StoreFollows(ctx, key, data); err != nil {
		// logging
	}

	return data, nil
}

// Followees 關注列表（分頁）
func (uc *FollowUc) Followees(ctx context.Context, req *model.SearchFolloweeCond) ([]*model.Follow, error) {
	cacheData, err := uc.cacheRepo.Followees(ctx, req)
	if err != nil {
		// logging
		return nil, err
	}

	if cacheData != nil {
		return cacheData, nil
	}

	data, err := uc.dbRepo.Followees(ctx, req)
	if err != nil {
		// logging
		return nil, err
	}

	key := fmt.Sprintf("%s:%s", model.CacheKey_Followees, req.UserId)
	if err := uc.cacheRepo.StoreFollows(ctx, key, data); err != nil {
		// logging
	}

	return data, nil
}

// Friends 好友列表（分頁）
func (uc *FollowUc) Friends(ctx context.Context, req *model.SearchFriendCond) ([]*model.Friend, error) {
	cacheData, err := uc.cacheRepo.Friends(ctx, req)
	if err != nil {
		// logging
		return nil, err
	}

	if cacheData != nil {
		return cacheData, nil
	}

	data, err := uc.dbRepo.Friends(ctx, req)
	if err != nil {
		// logging
		return nil, err
	}

	key := fmt.Sprintf("%s:%s", model.CacheKey_Friends, req.UserId)
	if err := uc.cacheRepo.StoreFriends(ctx, key, data); err != nil {
		// logging
	}

	return data, nil
}

// CheckAndSyncConsumerTime
// 如果傳入的時間比 cache 的早，則回傳 false
// 如果傳入的時間比 cache 的晚，則回傳 ture 並且將 cache 的時間更新為傳入的時間
func (uc *FollowUc) CheckAndSyncConsumerTime(ctx context.Context, t time.Time) bool {
	key := ""
	data, err := uc.cacheRepo.GetConsumerTime(ctx, key)
	if err != nil {
		// logging
		return false
	}

	if data.After(t) {
		return false
	}

	if err := uc.cacheRepo.UpsertConsumerTime(ctx, key, t); err != nil {
		return false
	}

	return true
}
