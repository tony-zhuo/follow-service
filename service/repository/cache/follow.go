package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tony-zhuo/follow-service/pkg/libs"
	pkgRedis "github.com/tony-zhuo/follow-service/pkg/redis"
	"github.com/tony-zhuo/follow-service/service/model"
	"sync"
	"time"
)

var (
	followRepo     *FollowRepo
	followRepoOnce sync.Once
)

type FollowRepo struct {
	cache *pkgRedis.RedisClient
}

func NewFollowCacheRepo(cache *pkgRedis.RedisClient) *FollowRepo {
	followRepoOnce.Do(func() {
		followRepo = &FollowRepo{
			cache: cache,
		}
	})
	return followRepo
}

// Follow 關注
func (repo *FollowRepo) Follow(ctx context.Context, req *model.Follow) error {
	key := ""
	_, err := repo.cache.ZAdd(ctx, key, &redis.Z{
		Score:  float64(req.CreatedAt.UnixNano()),
		Member: req,
	}).Result()

	return err
}

// UnFollow 取消關注
func (repo *FollowRepo) UnFollow(ctx context.Context, req *model.Follow) error {
	key := ""
	_, err := repo.cache.ZRem(ctx, key, &redis.Z{
		Score:  float64(req.CreatedAt.UnixNano()),
		Member: req,
	}).Result()

	return err
}

// Followers 粉絲清單（分頁）
func (repo *FollowRepo) Followers(ctx context.Context, cond *model.SearchFollowerCond) ([]*model.Follow, error) {
	key := fmt.Sprintf("%s:%s", model.CacheKey_Followers, cond.UserId)
	var redisCond *redis.ZRangeBy
	if cond.NextTimestamp == nil {
		redisCond = &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  int64(cond.Limit),
		}
	} else {
		redisCond = &redis.ZRangeBy{
			Min:    fmt.Sprintf("%d", *cond.NextTimestamp),
			Max:    "+inf",
			Offset: 0,
			Count:  int64(cond.Limit),
		}
	}
	cacheData, err := repo.cache.ZRangeByScoreWithScores(ctx, key, redisCond).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	res := make([]*model.Follow, len(cacheData))
	for i, data := range cacheData {
		tmp := data.Member.(model.Follow)
		res[i] = &tmp
	}

	return res, nil
}

// Followees 關注列表（分頁）
func (repo *FollowRepo) Followees(ctx context.Context, cond *model.SearchFolloweeCond) ([]*model.Follow, error) {
	key := fmt.Sprintf("%s:%s", model.CacheKey_Followees, cond.UserId)
	var redisCond *redis.ZRangeBy
	if cond.NextTimestamp == nil {
		redisCond = &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  int64(cond.Limit),
		}
	} else {
		redisCond = &redis.ZRangeBy{
			Min:    fmt.Sprintf("%d", *cond.NextTimestamp),
			Max:    "+inf",
			Offset: 0,
			Count:  int64(cond.Limit),
		}
	}
	cacheData, err := repo.cache.ZRangeByScoreWithScores(ctx, key, redisCond).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	res := make([]*model.Follow, len(cacheData))
	for i, data := range cacheData {
		tmp := data.Member.(model.Follow)
		res[i] = &tmp
	}

	return res, nil
}

// Friends 好友列表（分頁）
func (repo *FollowRepo) Friends(ctx context.Context, cond *model.SearchFriendCond) ([]*model.Friend, error) {
	key := fmt.Sprintf("%s:%s", model.CacheKey_Friends, cond.UserId)
	var redisCond *redis.ZRangeBy
	if cond.NextTimestamp == nil {
		redisCond = &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  int64(cond.Limit),
		}
	} else {
		redisCond = &redis.ZRangeBy{
			Min:    fmt.Sprintf("%d", *cond.NextTimestamp),
			Max:    "+inf",
			Offset: 0,
			Count:  int64(cond.Limit),
		}
	}
	cacheData, err := repo.cache.ZRangeByScoreWithScores(ctx, key, redisCond).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	res := make([]*model.Friend, len(cacheData))
	for i, data := range cacheData {
		tmp := data.Member.(model.Friend)
		res[i] = &tmp
	}

	return res, nil
}

// GetConsumerTime 取得 consumer 的 cache time
func (repo *FollowRepo) GetConsumerTime(ctx context.Context, key string) (*time.Time, error) {
	t, err := repo.cache.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	res, err := libs.StringConvertToTime(t)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpsertConsumerTime 更新或新增 consumer 的 cache time
func (repo *FollowRepo) UpsertConsumerTime(ctx context.Context, key string, value time.Time) error {
	res, err := repo.GetConsumerTime(ctx, key)
	if err != nil {
		return err
	}

	_, err = repo.cache.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		if res != nil {
			pipe.Del(ctx, key)
		}
		pipe.Set(ctx, key, value, -1)
		return nil
	})

	return err
}

func (repo *FollowRepo) StoreFollows(ctx context.Context, key string, data []*model.Follow) error {

	return nil
}

func (repo *FollowRepo) StoreFriends(ctx context.Context, key string, data []*model.Friend) error {

	return nil
}
