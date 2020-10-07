package service

import (
	"context"
	"errors"
	"fmt"
	"gostudy/microservice/gokit/user/pb"
)

type baseServer struct{}

func NewService() pb.UserServer {
	return &baseServer{}
}

// Login logic
func (s baseServer) Login(ctx context.Context, in *pb.LoginReq) (tok *pb.LoginRes, err error) {
	fmt.Println("调用 service Login 处理请求")
	if in.Username != "LeonardWang" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	tok = &pb.LoginRes{Token: "Test Token"}
	err = nil
	return
}
