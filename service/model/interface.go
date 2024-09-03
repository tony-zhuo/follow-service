package model

import (
	"context"
	"time"
)

type FollowUcInterface interface {
	// Follow 關注
	// 關注後會先存在 cache，並且丟到 kafka 進行後續處理
	Follow(ctx context.Context, req *FollowRequest) error

	// FollowInDB 關注
	FollowInDB(ctx context.Context, req *FollowRequest) error

	// UnFollow 取消關注
	// 取消關注後會先存在 cache，並且丟到 kafka 進行後續處理
	UnFollow(ctx context.Context, req *FollowRequest) error

	// UnFollowInDB 取消關注
	UnFollowInDB(ctx context.Context, req *FollowRequest) error

	// Followers 粉絲清單（分頁）
	Followers(ctx context.Context, req *SearchFollowerCond) ([]*Follow, error)

	// Followees 關注列表（分頁）
	Followees(ctx context.Context, req *SearchFolloweeCond) ([]*Follow, error)

	// Friends 好友列表（分頁）
	Friends(ctx context.Context, req *SearchFriendCond) ([]*Friend, error)

	// CheckAndSyncConsumerTime
	// 如果傳入的時間比 cache 的早，則回傳 false
	// 如果傳入的時間比 cache 的晚，則回傳 ture 並且將 cache 的時間更新為傳入的時間
	CheckAndSyncConsumerTime(ctx context.Context, t time.Time) bool
}

type FollowCacheRepoInterface interface {
	// Follow 關注
	Follow(ctx context.Context, req *Follow) error

	// UnFollow 取消關注
	UnFollow(ctx context.Context, req *Follow) error

	// Followers 粉絲清單（分頁）
	Followers(ctx context.Context, cond *SearchFollowerCond) ([]*Follow, error)

	// Followees 關注列表（分頁）
	Followees(ctx context.Context, cond *SearchFolloweeCond) ([]*Follow, error)

	// Friends 好友列表（分頁）
	Friends(ctx context.Context, cond *SearchFriendCond) ([]*Friend, error)

	// GetConsumerTime 取得 consumer 的 cache time
	GetConsumerTime(ctx context.Context, key string) (*time.Time, error)

	// UpsertConsumerTime 更新或新增 consumer 的 cache time
	UpsertConsumerTime(ctx context.Context, key string, value time.Time) error

	// StoreFollows 將 follow 列表存在 cache
	StoreFollows(ctx context.Context, key string, data []*Follow) error

	// StoreFriends 將 friend 列表存在 cache
	StoreFriends(ctx context.Context, key string, data []*Friend) error
}

type FollowDBRepoInterface interface {
	// Follow 關注
	// 當 follower_id 關注 followee_id 時，如 followee_id 已經有關注 follower_id
	// 則將 follower_id 跟 followee_id 互相加到 friend 中
	Follow(ctx context.Context, req *FollowRequest) error

	// UnFollow 取消關注
	// 當 follower_id 取消關注 followee_id 時，如 followee_id 已經有關注 follower_id
	// 將 follower_id 跟 followee_id 互相從 friend 中刪除
	UnFollow(ctx context.Context, req *FollowRequest) error

	// Followers 粉絲清單（分頁）
	// cond.Next 為上一個 resp 的最後一個 follower_id
	// return.[]string 為 follower_id
	Followers(ctx context.Context, cond *SearchFollowerCond) ([]*Follow, error)

	// Followees 關注列表（分頁）
	// cond.Next 為上一個 resp 的最後一個 followee_id
	// return.[]string 為 follower_id
	Followees(ctx context.Context, cond *SearchFolloweeCond) ([]*Follow, error)

	// Friends 好友列表（分頁）
	// cond.Next 為上一個 resp 的最後一個 friend_id
	// return.[]string 為 friend_id
	Friends(ctx context.Context, cond *SearchFriendCond) ([]*Friend, error)
}
