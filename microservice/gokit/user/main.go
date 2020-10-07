package main

import (
	"net"
	"os"

	"google.golang.org/grpc"

	"gostudy/microservice/gokit/user/endpoint"
	"gostudy/microservice/gokit/user/pb"
	"gostudy/microservice/gokit/user/service"
	"gostudy/microservice/gokit/user/transport"
)

func main() {
	ser := service.NewService()
	endpoints := endpoint.NewEndPointServer(ser)
	grpcServer := transport.NewGRPCServer(endpoints)

	grpcListener, err := net.Listen("tcp", ":8881")
	if err != nil {
		os.Exit(0)
	}
	gs := grpc.NewServer()
	pb.RegisterUserServer(gs, grpcServer)
	if err = gs.Serve(grpcListener); err != nil {
		os.Exit(0)
	}
}
