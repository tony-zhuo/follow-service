package db

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/tony-zhuo/follow-service/pkg/db"
	"github.com/tony-zhuo/follow-service/service/model"
	"sync"
)

var (
	followRepo     *FollowRepo
	followRepoOnce sync.Once
)

type FollowRepo struct {
	db *db.CassandraDB
}

func NewFollowDBRepo(db *db.CassandraDB) *FollowRepo {
	followRepoOnce.Do(func() {
		followRepo = &FollowRepo{
			db: db,
		}
	})
	return followRepo
}

// Follow 關注
// 當 follower_id 關注 followee_id 時，如 followee_id 已經有關注 follower_id
// 則將 follower_id 跟 followee_id 互相加到 friend 中
func (repo *FollowRepo) Follow(ctx context.Context, req *model.FollowRequest) error {
	followStmt := `INSERT INTO follow.follows (follower_id, followee_id, created_at) VALUES (?, ?, toTimestamp(now()))`
	friendStmt := `INSERT INTO follow.friends (user_id, friend_id, created_at)
		  VALUES (?, ?, toTimestamp(now()))
		  IF EXISTS (
			SELECT * FROM follow.follows WHERE follower_id = ? AND followee_id = ?
		  );`
	batch := repo.db.Instance.NewBatch(gocql.LoggedBatch).WithContext(ctx)
	batch.Query(followStmt, req.FollowerId, req.FolloweeId)
	batch.Query(friendStmt, req.FollowerId, req.FolloweeId, req.FolloweeId, req.FollowerId)
	batch.Query(friendStmt, req.FolloweeId, req.FollowerId, req.FollowerId, req.FolloweeId)
	return repo.db.Instance.ExecuteBatch(batch)
}

// UnFollow 取消關注
// 當 follower_id 取消關注 followee_id 時，如 followee_id 已經有關注 follower_id
// 將 follower_id 跟 followee_id 互相從 friend 中刪除
func (repo *FollowRepo) UnFollow(ctx context.Context, req *model.FollowRequest) error {
	followStmt := `DELETE FROM follow.follows WHERE follower_id = ? AND followee_id = ?`
	friendStmt := `DELETE FROM follow.friends WHERE user_id = ? AND friend_id = ?
		  IF EXISTS (
			SELECT * FROM follow.follows WHERE follower_id = ? AND followee_id = ?
		  );`

	batch := repo.db.Instance.NewBatch(gocql.LoggedBatch).WithContext(ctx)
	batch.Query(followStmt, req.FollowerId, req.FolloweeId)
	batch.Query(friendStmt, req.FollowerId, req.FolloweeId, req.FolloweeId, req.FollowerId)
	batch.Query(friendStmt, req.FolloweeId, req.FollowerId, req.FollowerId, req.FolloweeId)
	return repo.db.Instance.ExecuteBatch(batch)
}

// Followers 粉絲清單（分頁）
// cond.Next 為上一個 resp 的最後一個 follower_id
// return.[]string 為 follower_id
func (repo *FollowRepo) Followers(ctx context.Context, cond *model.SearchFollowerCond) ([]*model.Follow, error) {
	var res []*model.Follow
	var err error
	stmt := `SELECT * FROM follows
			WHERE followee_id = ?
			AND token(follower_id) > token(?)
			ORDER BY created_at ASC
			LIMIT ?;`
	if cond.NextUserId != nil {
		err = repo.db.Instance.Query(stmt, cond.UserId, cond.NextUserId, cond.Limit).WithContext(ctx).Scan(res)
	} else {
		err = repo.db.Instance.Query(stmt, cond.UserId, 0, cond.Limit).WithContext(ctx).Scan(res)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Followees 關注列表（分頁）
// cond.Next 為上一個 resp 的最後一個 followee_id
// return.[]string 為 follower_id
func (repo *FollowRepo) Followees(ctx context.Context, cond *model.SearchFolloweeCond) ([]*model.Follow, error) {
	var res []*model.Follow
	var err error
	stmt := `SELECT followee_id FROM follows
			WHERE follower_id = ?
			AND token(followee_id) > token(?)
			ORDER BY created_at ASC
			LIMIT ?;`

	if cond.NextUserId != nil {
		err = repo.db.Instance.Query(stmt, cond.UserId, cond.NextUserId, cond.Limit).WithContext(ctx).Scan(res)
	} else {
		err = repo.db.Instance.Query(stmt, cond.UserId, 0, cond.Limit).WithContext(ctx).Scan(res)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Friends 好友列表（分頁）
// cond.Next 為上一個 resp 的最後一個 friend_id
// return.[]string 為 friend_id
func (repo *FollowRepo) Friends(ctx context.Context, cond *model.SearchFriendCond) ([]*model.Friend, error) {
	var res []*model.Friend
	var err error
	stmt := `SELECT friend_id FROM friends
			WHERE user_id = ?
			AND token(friend_id) > token(?)
			ORDER BY created_at ASC
			LIMIT ?;`

	if cond.NextFriendId != nil {
		err = repo.db.Instance.Query(stmt, cond.UserId, cond.NextFriendId, cond.Limit).WithContext(ctx).Scan(res)
	} else {
		err = repo.db.Instance.Query(stmt, cond.UserId, 0, cond.Limit).WithContext(ctx).Scan(res)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
