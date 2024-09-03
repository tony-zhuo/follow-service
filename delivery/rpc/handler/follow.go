package handler

import (
	"context"
	proto "github.com/tony-zhuo/follow-service/protos/data"
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

func (grpc *FollowGrpc) Follow(ctx context.Context, req *proto.FollowReq) (*proto.FollowResp, error) {
	err := grpc.uc.Follow(ctx, &model.FollowRequest{
		FollowerId: req.GetFollowerId(),
		FolloweeId: req.GetFolloweeId(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.FollowResp{
		Resp: &proto.CommonResp{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (grpc *FollowGrpc) UnFollow(ctx context.Context, req *proto.FollowReq) (*proto.FollowResp, error) {
	err := grpc.uc.UnFollow(ctx, &model.FollowRequest{
		FollowerId: req.GetFollowerId(),
		FolloweeId: req.GetFolloweeId(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.FollowResp{
		Resp: &proto.CommonResp{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (grpc *FollowGrpc) Followers(ctx context.Context, req *proto.FollowersReq) (*proto.FollowersResp, error) {
	data, err := grpc.uc.Followers(ctx, &model.SearchFollowerCond{
		UserId:        req.GetUserId(),
		NextUserId:    nil,
		NextTimestamp: nil,
		Limit:         int(req.GetLimit()),
	})
	if err != nil {
		return nil, err
	}

	// TODO: data convert to proto.Follow

	return &proto.FollowersResp{
		Resp: &proto.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		Data: make([]*proto.Follow, 0),
	}, nil
}

func (grpc *FollowGrpc) Followees(ctx context.Context, req *proto.FolloweesReq) (*proto.FolloweesResp, error) {
	data, err := grpc.uc.Followees(ctx, &model.SearchFolloweeCond{
		UserId:        req.GetUserId(),
		NextUserId:    nil,
		NextTimestamp: nil,
		Limit:         int(req.GetLimit()),
	})
	if err != nil {
		return nil, err
	}

	// TODO: data convert to proto.Follow

	return &proto.FolloweesResp{
		Resp: &proto.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		Data: make([]*proto.Follow, 0),
	}, nil
}

func (grpc *FollowGrpc) Friends(ctx context.Context, req *proto.FriendsReq) (*proto.FriendsResp, error) {
	data, err := grpc.uc.Friends(ctx, &model.SearchFriendCond{
		UserId:        req.GetUserId(),
		NextFriendId:  nil,
		NextTimestamp: nil,
		Limit:         int(req.GetLimit()),
	})
	if err != nil {
		return nil, err
	}

	// TODO: data convert to proto.Friend

	return &proto.FriendsResp{
		Resp: &proto.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		Data: make([]*proto.Friend, 0),
	}, nil
}
