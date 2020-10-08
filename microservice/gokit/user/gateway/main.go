package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc"

	"gostudy/microservice/gokit/user/pb"
)

func main() {

	ctx := context.Background()
	//Etcd客户端
	client, _ := etcdv3.NewClient(ctx, []string{"127.0.0.1:2379"}, etcdv3.ClientOptions{})

	//服务实例
	instancer, _ := etcdv3.NewInstancer(client, "svc.user", log.NewNopLogger())

	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, reqFactory, log.NewNopLogger()) //reqFactory自定义的函数，主要用于端点层（endpoint）接受并显示数据
	//创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)

	reqEndPoint := lb.Retry(3, 3*time.Second, balancer)

	//现在我们可以通过 endPoint 发起请求了
	req := &pb.LoginReq{
		Username: "LeonardWang",
		Password: "123456",
	}
	if res, err := reqEndPoint(ctx, req); err != nil {
		panic(err)
	} else {
		fmt.Println(res, err)
	}

	return
}

//通过传入的 实例地址  创建对应的请求endPoint
func reqFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	fmt.Println("instanceAddr:", instanceAddr)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("请求服务: ", instanceAddr)
		conn, err := grpc.Dial(instanceAddr, grpc.WithInsecure())
		fmt.Println("new conn!")
		if err != nil {
			fmt.Println(err)
			panic("connect error")
		}
		defer conn.Close()
		svr := pb.NewUserClient(conn)
		switch t := request.(type) {
		case *pb.LoginReq:
			return svr.Login(ctx, t)
		default:
			return nil, errors.New("Unknown Type")
		}
	}, nil, nil
}
