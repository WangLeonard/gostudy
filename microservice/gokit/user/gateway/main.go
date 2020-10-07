package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gostudy/microservice/gokit/user/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	svr := pb.NewUserClient(conn)
	tok, err := svr.Login(context.Background(), &pb.LoginReq{
		Username: "LeonardWang",
		Password: "123456",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tok.Token)
}
