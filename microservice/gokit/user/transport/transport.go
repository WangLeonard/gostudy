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
	//options := []grpctransport.ServerOption{
	//	grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
	//		ctx = context.WithValue(ctx, v5_service.ContextReqUUid, md.Get(v5_service.ContextReqUUid))
	//		return ctx
	//	}),
	//	//grpctransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
	//}
	return &grpcServer{login: grpctransport.NewServer(
		endpoint.LoginEndPoint,
		RequestGrpcLogin,
		ResponseGrpcLogin,
		//options...,
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
	req := grpcReq.(*pb.LoginReq)
	return &pb.LoginReq{Username: req.GetUsername(), Password: req.GetPassword()}, nil
}

func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginRes)
	return &pb.LoginRes{Token: resp.Token}, nil
}
