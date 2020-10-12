package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/sd/etcdv3"
)

var etcdAddr = "127.0.0.1:2379"

func main() {
	ctx := context.Background()
	//Etcd客户端
	client, _ := etcdv3.NewClient(ctx, []string{etcdAddr}, etcdv3.ClientOptions{})

	r := gin.Default()
	ApiGroup := r.Group("")

	// register user client.
	ConnectUserService(ctx, client, ApiGroup)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
