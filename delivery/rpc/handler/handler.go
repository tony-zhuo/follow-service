package handler

import (
	"github.com/tony-zhuo/follow-service/delivery/rpc/init"
	"github.com/tony-zhuo/follow-service/pkg/grpclib"
	proto "github.com/tony-zhuo/follow-service/protos/data"
	"google.golang.org/grpc"
)

func Start() {
	init.InitConf()

	grpcServer := grpclib.NewServer()
	InitServer(grpcServer)
}

func InitServer(grpcServer *grpc.Server) {
	proto.RegisterFollowServiceServer(grpcServer, NewNotificationGrpc())
}
