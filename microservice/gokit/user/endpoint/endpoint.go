package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gostudy/microservice/gokit/user/pb"
	"gostudy/microservice/gokit/user/service"
)

type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc service.Service) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
	}
	return EndPointServer{LoginEndPoint: loginEndPoint}
}

func (s EndPointServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginRes), nil
}

func MakeLoginEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LoginReq)
		return s.Login(ctx, req)
	}
}
