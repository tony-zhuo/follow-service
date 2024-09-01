package model

import "time"

type Follow struct {
	FollowerId string    `json:"follower_id"` // 關注者
	FolloweeId string    `json:"followee_id"` // 被關注者
	CreatedAt  time.Time `json:"created_at"`  // 關注時間
}

type Friend struct {
	UserId    string    `json:"user_id"`    // user id
	FriendId  string    `json:"friend_id"`  // 與 user 有互相關注的人
	CreatedAt time.Time `json:"created_at"` // 成為好友的時間
}
