// package transport trans endpoint to grpc
// TODO: Automatic generated the file by pb.

package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"gostudy/microservice/gokit/user/endpoint"
	userpb "gostudy/microservice/gokit/user/pb"
)

type grpcServer struct {
	regist grpctransport.Handler
	login  grpctransport.Handler
}

func NewGRPCServer(endpoint endpoint.EndPointServer) userpb.UserServer {
	return &grpcServer{
		regist: grpctransport.NewServer(
			endpoint.RegistEndPoint,
			RequestGrpcRegist,
			ResponseGrpcRegist,
		),
		login: grpctransport.NewServer(
			endpoint.LoginEndPoint,
			RequestGrpcLogin,
			ResponseGrpcLogin,
		)}
}

func (s *grpcServer) Regist(ctx context.Context, req *userpb.RegistReq) (*userpb.RegistResp, error) {
	_, rep, err := s.regist.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*userpb.RegistResp), nil
}

func RequestGrpcRegist(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func ResponseGrpcRegist(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func (s *grpcServer) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*userpb.LoginResp), nil
}

func RequestGrpcLogin(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
