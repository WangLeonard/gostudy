package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"gostudy/microservice/gokit/user/endpoint"
	"gostudy/microservice/gokit/user/pb"
	"gostudy/microservice/gokit/user/service"
	"gostudy/microservice/gokit/user/transport"
	"net"
	"os"
)

func main() {
	server := service.NewService()
	endpoints := endpoint.NewEndPointServer(server)
	grpcServer := transport.NewGRPCServer(endpoints)
	grpcListener, err := net.Listen("tcp", ":8881")
	if err != nil {
		//utils.GetLogger().Warn("Listen", zap.Error(err))
		os.Exit(0)
	}
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(baseServer, grpcServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		//utils.GetLogger().Warn("Serve", zap.Error(err))
		os.Exit(0)
	}
}
