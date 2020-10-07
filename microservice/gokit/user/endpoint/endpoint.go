package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gostudy/microservice/gokit/user/pb"
)

type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
	// Add new EndPoint here.
}

func NewEndPointServer(svc pb.UserServer) EndPointServer {
	return EndPointServer{LoginEndPoint: makeLoginEndPoint(svc)}
}

func (s EndPointServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginRes), nil
}

func makeLoginEndPoint(s pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LoginReq)
		return s.Login(ctx, req)
	}
}
