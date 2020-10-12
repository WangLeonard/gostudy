package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"

	"gostudy/microservice/gokit/demo/rpc/user/endpoint"
	userpb "gostudy/microservice/gokit/demo/rpc/user/pb"
	"gostudy/microservice/gokit/demo/rpc/user/service"
	"gostudy/microservice/gokit/demo/rpc/user/transport"
)

func main() {
	// TODO: use config file.
	var (
		etcdAddrs = []string{"127.0.0.1:2379"}
		serName   = "svc.user"
		grpcAddr  = "127.0.0.1:8881"
		ttl       = 5 * time.Second
	)

	// 初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		fmt.Println(err)
		return
	}
	Registar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   serName,
		Value: grpcAddr,
	}, log.NewNopLogger())

	// 注册 etcd
	Registar.Register()

	// 初始化grpc处理逻辑
	ser := service.NewService()
	endpoints := endpoint.NewUserEndPointServer(ser)
	grpcServer := transport.NewGRPCServer(endpoints)

	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		os.Exit(0)
	}
	// 建立 grpc 服务端
	gs := grpc.NewServer()
	// 注册 grpc 服务端
	userpb.RegisterUserServer(gs, grpcServer)
	// 启动监听
	if err = gs.Serve(grpcListener); err != nil {
		os.Exit(0)
	}
}
