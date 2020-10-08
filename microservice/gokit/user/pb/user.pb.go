// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_56ad8ed12319a0be, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (dst *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(dst, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRes struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_56ad8ed12319a0be, []int{1}
}
func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (dst *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(dst, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "pb.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "pb.LoginRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/pb.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_56ad8ed12319a0be) }

var fileDescriptor_user_56ad8ed12319a0be = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe2, 0xe2, 0xf0, 0xc9,
	0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x14, 0x92, 0xe2, 0xe2, 0x08, 0x2d, 0x4e, 0x2d, 0xca, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x41, 0x72, 0x01, 0x89, 0xc5,
	0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x10, 0x39, 0x18, 0x5f, 0x49, 0x01, 0x6e, 0x46, 0xb1,
	0x90, 0x08, 0x17, 0x6b, 0x48, 0x7e, 0x76, 0x6a, 0x1e, 0xd4, 0x00, 0x08, 0xc7, 0x48, 0x97, 0x8b,
	0x05, 0x64, 0x92, 0x90, 0x2a, 0x17, 0x2b, 0x58, 0xa5, 0x10, 0x8f, 0x5e, 0x41, 0x92, 0x1e, 0xcc,
	0x62, 0x29, 0x64, 0x5e, 0xb1, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x7d, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x98, 0x8f, 0x87, 0x9d, 0xad, 0x00, 0x00, 0x00,
}
