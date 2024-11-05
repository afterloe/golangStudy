// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: user_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServerClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userServerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServerClient(cc grpc.ClientConnInterface) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/Logic.UserServer/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
// All implementations must embed UnimplementedUserServerServer
// for forward compatibility
type UserServerServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedUserServerServer()
}

// UnimplementedUserServerServer must be embedded to have forward compatible implementations.
type UnimplementedUserServerServer struct {
}

func (UnimplementedUserServerServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServerServer) mustEmbedUnimplementedUserServerServer() {}

// UnsafeUserServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServerServer will
// result in compilation errors.
type UnsafeUserServerServer interface {
	mustEmbedUnimplementedUserServerServer()
}

func RegisterUserServerServer(s grpc.ServiceRegistrar, srv UserServerServer) {
	s.RegisterService(&UserServer_ServiceDesc, srv)
}

func _UserServer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Logic.UserServer/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServer_ServiceDesc is the grpc.ServiceDesc for UserServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Logic.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserServer_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
