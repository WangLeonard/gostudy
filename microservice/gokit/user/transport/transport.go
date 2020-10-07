package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"gostudy/microservice/gokit/user/endpoint"
	"gostudy/microservice/gokit/user/pb"
)

type grpcServer struct {
	login grpctransport.Handler
}

func NewGRPCServer(endpoint endpoint.EndPointServer) pb.UserServer {
	return &grpcServer{login: grpctransport.NewServer(
		endpoint.LoginEndPoint,
		RequestGrpcLogin,
		ResponseGrpcLogin,
	)}
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginRes), nil
}

func RequestGrpcLogin(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
