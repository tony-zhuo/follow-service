package grpclib

import (
	"go.elastic.co/apm/module/apmgrpc/v2"
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
		),
		grpc.ChainStreamInterceptor(
			apmgrpc.NewStreamServerInterceptor(),
		),
	)
}
