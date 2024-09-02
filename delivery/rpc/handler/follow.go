package handler

import (
	"context"
	"github.com/tony-zhuo/follow-service/service/model"
	"sync"
)

var followGrpc *FollowGrpc
var followGrpcOnce sync.Once

func NewNotificationGrpc(uc model.FollowUcInterface) *FollowGrpc {
	followGrpcOnce.Do(func() {
		followGrpc = &FollowGrpc{
			uc: uc,
		}
	})
	return followGrpc
}

type FollowGrpc struct {
	uc model.FollowUcInterface
}

func (grpc *FollowGrpc) Follow(ctx context.Context) {}

func (grpc *FollowGrpc) UnFollow(ctx context.Context) {}

func (grpc *FollowGrpc) Followers(ctx context.Context) {}

func (grpc *FollowGrpc) Followees(ctx context.Context) {}

func (grpc *FollowGrpc) Friends(ctx context.Context) {}
