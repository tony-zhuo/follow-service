package model

type Action string

const (
	Action_Follow   Action = "follow"
	Action_UnFollow Action = "unfollow"
)

func (a Action) ToString() string {
	return string(a)
}

type CacheKey string

const (
	CacheKey_Followers CacheKey = "followers"
	CacheKey_Followees CacheKey = "followees"
	CacheKey_Friends   CacheKey = "friends"
)
