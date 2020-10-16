package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
)

func ConnectFileService(ctx context.Context, etcdClient etcdv3.Client, ginGroup *gin.RouterGroup) {
	//服务实例
	instancer, _ := etcdv3.NewInstancer(etcdClient, "svc.file", log.NewNopLogger())
	//fmt.Println(instancer.cache.State().Instances) its addr

	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, fileReqFactory, log.NewNopLogger()) //reqFactory自定义的函数，主要用于端点层（endpoint）接受并显示数据
	//创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	reqEndPoint := lb.Retry(1, 3*time.Second, balancer)

	uploadHandle := func(c *gin.Context) {
		reqEndPoint(ctx, c)
		//res, err := reqEndPoint(ctx, c)
		//fmt.Println("Handler", res, err)
	}
	fileGroup := ginGroup.Group("file")

	fileGroup.POST("/upload", uploadHandle)
}

// 通过传入的 实例地址  创建对应的请求endPoint
func fileReqFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ginCtx := request.(*gin.Context)
		simpleHostProxy := httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = "http"
				req.URL.Host = instanceAddr
				req.Host = instanceAddr
			},
		}

		// 转发
		simpleHostProxy.ServeHTTP(ginCtx.Writer, ginCtx.Request)
		//fmt.Println("fileReqFactory:", ginCtx.Request)
		return ginCtx.Request, nil
	}, nil, nil
}
