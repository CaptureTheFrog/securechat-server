// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ClientServerComms.proto

package grpc

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

// ClientServerCommsClient is the client API for ClientServerComms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServerCommsClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignUpChallengeResponse(ctx context.Context, in *SignUpChallengeResponseRequest, opts ...grpc.CallOption) (*SignUpChallengeResponseResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error)
}

type clientServerCommsClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServerCommsClient(cc grpc.ClientConnInterface) ClientServerCommsClient {
	return &clientServerCommsClient{cc}
}

func (c *clientServerCommsClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/main.ClientServerComms/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServerCommsClient) SignUpChallengeResponse(ctx context.Context, in *SignUpChallengeResponseRequest, opts ...grpc.CallOption) (*SignUpChallengeResponseResponse, error) {
	out := new(SignUpChallengeResponseResponse)
	err := c.cc.Invoke(ctx, "/main.ClientServerComms/SignUpChallengeResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServerCommsClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/main.ClientServerComms/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServerCommsClient) FindUser(ctx context.Context, in *FindUserRequest, opts ...grpc.CallOption) (*FindUserResponse, error) {
	out := new(FindUserResponse)
	err := c.cc.Invoke(ctx, "/main.ClientServerComms/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServerCommsServer is the server API for ClientServerComms service.
// All implementations must embed UnimplementedClientServerCommsServer
// for forward compatibility
type ClientServerCommsServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignUpChallengeResponse(context.Context, *SignUpChallengeResponseRequest) (*SignUpChallengeResponseResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error)
	mustEmbedUnimplementedClientServerCommsServer()
}

// UnimplementedClientServerCommsServer must be embedded to have forward compatible implementations.
type UnimplementedClientServerCommsServer struct {
}

func (UnimplementedClientServerCommsServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedClientServerCommsServer) SignUpChallengeResponse(context.Context, *SignUpChallengeResponseRequest) (*SignUpChallengeResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpChallengeResponse not implemented")
}
func (UnimplementedClientServerCommsServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedClientServerCommsServer) FindUser(context.Context, *FindUserRequest) (*FindUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedClientServerCommsServer) mustEmbedUnimplementedClientServerCommsServer() {}

// UnsafeClientServerCommsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServerCommsServer will
// result in compilation errors.
type UnsafeClientServerCommsServer interface {
	mustEmbedUnimplementedClientServerCommsServer()
}

func RegisterClientServerCommsServer(s grpc.ServiceRegistrar, srv ClientServerCommsServer) {
	s.RegisterService(&ClientServerComms_ServiceDesc, srv)
}

func _ClientServerComms_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServerCommsServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ClientServerComms/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServerCommsServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientServerComms_SignUpChallengeResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpChallengeResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServerCommsServer).SignUpChallengeResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ClientServerComms/SignUpChallengeResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServerCommsServer).SignUpChallengeResponse(ctx, req.(*SignUpChallengeResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientServerComms_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServerCommsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ClientServerComms/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServerCommsServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientServerComms_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServerCommsServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ClientServerComms/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServerCommsServer).FindUser(ctx, req.(*FindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientServerComms_ServiceDesc is the grpc.ServiceDesc for ClientServerComms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientServerComms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ClientServerComms",
	HandlerType: (*ClientServerCommsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _ClientServerComms_SignUp_Handler,
		},
		{
			MethodName: "SignUpChallengeResponse",
			Handler:    _ClientServerComms_SignUpChallengeResponse_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ClientServerComms_Login_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _ClientServerComms_FindUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ClientServerComms.proto",
}
