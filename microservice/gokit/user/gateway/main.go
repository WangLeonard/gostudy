package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc"

	userpb "gostudy/microservice/gokit/user/pb"
)

var etcdAddr = "127.0.0.1:2379"

func main() {

	ctx := context.Background()
	//Etcd客户端
	client, _ := etcdv3.NewClient(ctx, []string{etcdAddr}, etcdv3.ClientOptions{})

	//服务实例
	instancer, _ := etcdv3.NewInstancer(client, "svc.user", log.NewNopLogger())

	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, reqFactory, log.NewNopLogger()) //reqFactory自定义的函数，主要用于端点层（endpoint）接受并显示数据
	//创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)

	reqEndPoint := lb.Retry(1, 3*time.Second, balancer)

	registHandle := func(c *gin.Context) {
		var registReq = &userpb.RegistReq{}
		c.ShouldBindJSON(registReq)
		fmt.Println("registReq:", registReq)

		if res, err := reqEndPoint(ctx, registReq); err != nil {
			fmt.Println("err:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			fmt.Println("message:", res)
			c.JSON(http.StatusOK, gin.H{
				"message": res.(*userpb.RegistResp).Message,
			})
		}
	}

	loginHandle := func(c *gin.Context) {
		var loginReq = &userpb.LoginReq{}
		c.ShouldBindJSON(loginReq)

		if res, err := reqEndPoint(ctx, loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": res.(*userpb.LoginResp).Token,
			})
		}
	}

	r := gin.Default()
	ApiGroup := r.Group("")
	ApiGroup.GET("/regist", registHandle)
	ApiGroup.GET("/login", loginHandle)

	if err := r.Run(); err != nil {
		panic(err)
	}
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
		svr := userpb.NewUserClient(conn)
		switch t := request.(type) {
		case *userpb.RegistReq:
			return svr.Regist(ctx, t)
		case *userpb.LoginReq:
			return svr.Login(ctx, t)
		default:
			return nil, errors.New("Unknown Type")
		}
	}, nil, nil
}
