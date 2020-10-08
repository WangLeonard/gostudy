package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	userpb "gostudy/microservice/gokit/user/pb"
)

type EndPointServer struct {
	RegistEndPoint endpoint.Endpoint
	LoginEndPoint  endpoint.Endpoint
}

func NewUserEndPointServer(svc userpb.UserServer) EndPointServer {
	return EndPointServer{
		RegistEndPoint: makeRegistEndPoint(svc),
		LoginEndPoint:  makeLoginEndPoint(svc),
	}
}

func (s EndPointServer) Regist(ctx context.Context, in *userpb.RegistReq) (*userpb.RegistResp, error) {
	res, err := s.RegistEndPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*userpb.RegistResp), nil
}

func (s EndPointServer) Login(ctx context.Context, in *userpb.LoginReq) (*userpb.LoginResp, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*userpb.LoginResp), nil
}

func makeRegistEndPoint(s userpb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*userpb.RegistReq)
		return s.Regist(ctx, req)
	}
}

func makeLoginEndPoint(s userpb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*userpb.LoginReq)
		return s.Login(ctx, req)
	}
}
