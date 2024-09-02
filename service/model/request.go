package model

type FollowRequest struct {
	FollowerId string `json:"follower_id"`
	FolloweeId string `json:"followee_id"`
}

type SearchFollowerCond struct {
	UserId        string  `json:"user_id"`
	NextUserId    *string `json:"next_user_id"`
	NextTimestamp *int64  `json:"next_timestamp"`
	Limit         int     `json:"limit"`
}

type SearchFolloweeCond struct {
	UserId        string  `json:"user_id"`
	NextUserId    *string `json:"next_user_id"`
	NextTimestamp *int64  `json:"next_timestamp"`
	Limit         int     `json:"limit"`
}

type SearchFriendCond struct {
	UserId        string  `json:"user_id"`
	NextFriendId  *string `json:"next_friend_id"`
	NextTimestamp *int64  `json:"next_timestamp"`
	Limit         int     `json:"limit"`
}
