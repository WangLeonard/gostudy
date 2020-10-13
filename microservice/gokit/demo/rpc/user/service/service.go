// package service implement user logic
// TODO: Automatic generated part content by pb(except logic implement).

package service

import (
	"context"
	"errors"
	"fmt"

	"gostudy/microservice/gokit/demo/rpc/user/model"
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
	if _, err := model.Find(in.Username); err == nil {
		return nil, errors.New("用户已注册")
	}

	u := &model.RegisterStruct{Username: in.Username, Password: in.Password}
	if err := model.Create(u); err == nil {
		return &userpb.RegistResp{Message: "Ok"}, nil
	} else {
		return nil, errors.New("数据库插入失败")
	}
}

// Login logic
func (s baseServer) Login(ctx context.Context, in *userpb.LoginReq) (tok *userpb.LoginResp, err error) {
	fmt.Println("调用 service Login 处理请求")
	if password, ok := userDate[in.Username]; ok && password == in.Password {
		return &userpb.LoginResp{Token: "Test Token"}, nil
	}
	if u, err := model.Find(in.Username); err == nil {
		if in.Password == u.Password {
			return &userpb.LoginResp{Token: "Test Token"}, nil
		} else {
			return nil, errors.New("用户密码错误")
		}
	} else {
		return nil, err
	}
}
