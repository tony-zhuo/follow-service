package handler

import (
	"github.com/tony-zhuo/follow-service/delivery/rpc/init"
	"github.com/tony-zhuo/follow-service/pkg/grpclib"
	"google.golang.org/grpc"
)

func Start() {
	init.InitConf()

	grpcServer := grpclib.NewServer()
	InitServer(grpcServer)
}

func InitServer(grpcServer *grpc.Server) {
	conf := init.InitConf()

}
