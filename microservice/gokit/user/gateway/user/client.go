package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	userendpoint "gostudy/microservice/gokit/user/endpoint"
	"gostudy/microservice/gokit/user/pb"
	"gostudy/microservice/gokit/user/service"
)

func NewGRPCClient(conn *grpc.ClientConn) service.Service {
	//options := []grpctransport.ClientOption{
	//	grpctransport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
	//		UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
	//		log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
	//		md.Set(v5_service.ContextReqUUid, UUID)
	//		ctx = metadata.NewOutgoingContext(context.Background(), *md)
	//		return ctx
	//	}),
	//}
	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"pb.User",
			"Login",
			RequestLogin,
			ResponseLogin,
			pb.LoginRes{}).Endpoint()
	}
	return userendpoint.EndPointServer{
		LoginEndPoint: loginEndpoint,
	}
}

func RequestLogin(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginReq)
	return &pb.LoginReq{Username: req.Username, Password: req.Password}, nil
}

func ResponseLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginRes)
	return &pb.LoginRes{Token: resp.Token}, nil
}
