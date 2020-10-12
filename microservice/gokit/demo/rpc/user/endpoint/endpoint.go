// package endpoint trans logic to endpoint
// TODO: Automatic generated the file by pb (except middleware).

package endpoint

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	userpb "gostudy/microservice/gokit/demo/rpc/user/pb"
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

func middleware(s string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			fmt.Println(s, "in")
			defer fmt.Println(s, "out")
			return next(ctx, request)
		}
	}
}

func makeRegistEndPoint(s userpb.UserServer) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*userpb.RegistReq)
		return s.Regist(ctx, req)
	}
	return endpoint.Chain(
		middleware("Regist"),
	)(f)
}

func makeLoginEndPoint(s userpb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*userpb.LoginReq)
		return s.Login(ctx, req)
	}
}
