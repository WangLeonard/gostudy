// package service implement user logic
// TODO: Automatic generated part content by pb(except logic implement).

package service

import (
	"context"
	"errors"
	"fmt"

	userpb "gostudy/microservice/gokit/demo/rpc/user/pb"
)

var userDate = make(map[string]string)

type baseServer struct{}

func NewService() userpb.UserServer {
	return &baseServer{}
}

// Regist logic
func (s baseServer) Regist(ctx context.Context, in *userpb.RegistReq) (tok *userpb.RegistResp, err error) {
	fmt.Println("调用 service Regist 处理请求")
	if _, ok := userDate[in.Username]; !ok {
		userDate[in.Username] = in.Password
		return &userpb.RegistResp{Message: "Ok"}, nil
	}
	return nil, errors.New("用户已注册")
}

// Login logic
func (s baseServer) Login(ctx context.Context, in *userpb.LoginReq) (tok *userpb.LoginResp, err error) {
	fmt.Println("调用 service Login 处理请求")
	if password, ok := userDate[in.Username]; ok && password == in.Password {
		return &userpb.LoginResp{Token: "Test Token"}, nil
	}
	return nil, errors.New("用户信息错误")
}
