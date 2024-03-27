// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: ServerComm.proto

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

// ServerCommsClient is the client API for ServerComms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerCommsClient interface {
	GetSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
	GetPredecessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
	ChangeSuccessor(ctx context.Context, in *ChangeSuccessor, opts ...grpc.CallOption) (*Response, error)
	ChangePredecessor(ctx context.Context, in *ChangePredecessor, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Record, error)
	Put(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Response, error)
}

type serverCommsClient struct {
	cc grpc.ClientConnInterface
}

func NewServerCommsClient(cc grpc.ClientConnInterface) ServerCommsClient {
	return &serverCommsClient{cc}
}

func (c *serverCommsClient) GetSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/main.ServerComms/getSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverCommsClient) GetPredecessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/main.ServerComms/getPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverCommsClient) ChangeSuccessor(ctx context.Context, in *ChangeSuccessor, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.ServerComms/changeSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverCommsClient) ChangePredecessor(ctx context.Context, in *ChangePredecessor, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.ServerComms/changePredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverCommsClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/main.ServerComms/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverCommsClient) Put(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.ServerComms/put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerCommsServer is the server API for ServerComms service.
// All implementations must embed UnimplementedServerCommsServer
// for forward compatibility
type ServerCommsServer interface {
	GetSuccessor(context.Context, *ID) (*ID, error)
	GetPredecessor(context.Context, *ID) (*ID, error)
	ChangeSuccessor(context.Context, *ChangeSuccessor) (*Response, error)
	ChangePredecessor(context.Context, *ChangePredecessor) (*Response, error)
	Get(context.Context, *ID) (*Record, error)
	Put(context.Context, *Record) (*Response, error)
	mustEmbedUnimplementedServerCommsServer()
}

// UnimplementedServerCommsServer must be embedded to have forward compatible implementations.
type UnimplementedServerCommsServer struct {
}

func (UnimplementedServerCommsServer) GetSuccessor(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuccessor not implemented")
}
func (UnimplementedServerCommsServer) GetPredecessor(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessor not implemented")
}
func (UnimplementedServerCommsServer) ChangeSuccessor(context.Context, *ChangeSuccessor) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSuccessor not implemented")
}
func (UnimplementedServerCommsServer) ChangePredecessor(context.Context, *ChangePredecessor) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePredecessor not implemented")
}
func (UnimplementedServerCommsServer) Get(context.Context, *ID) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServerCommsServer) Put(context.Context, *Record) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedServerCommsServer) mustEmbedUnimplementedServerCommsServer() {}

// UnsafeServerCommsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerCommsServer will
// result in compilation errors.
type UnsafeServerCommsServer interface {
	mustEmbedUnimplementedServerCommsServer()
}

func RegisterServerCommsServer(s grpc.ServiceRegistrar, srv ServerCommsServer) {
	s.RegisterService(&ServerComms_ServiceDesc, srv)
}

func _ServerComms_GetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCommsServer).GetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerComms/getSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCommsServer).GetSuccessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerComms_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCommsServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerComms/getPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCommsServer).GetPredecessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerComms_ChangeSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSuccessor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCommsServer).ChangeSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerComms/changeSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCommsServer).ChangeSuccessor(ctx, req.(*ChangeSuccessor))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerComms_ChangePredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePredecessor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCommsServer).ChangePredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerComms/changePredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCommsServer).ChangePredecessor(ctx, req.(*ChangePredecessor))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerComms_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCommsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerComms/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCommsServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerComms_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCommsServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerComms/put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCommsServer).Put(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerComms_ServiceDesc is the grpc.ServiceDesc for ServerComms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerComms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ServerComms",
	HandlerType: (*ServerCommsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSuccessor",
			Handler:    _ServerComms_GetSuccessor_Handler,
		},
		{
			MethodName: "getPredecessor",
			Handler:    _ServerComms_GetPredecessor_Handler,
		},
		{
			MethodName: "changeSuccessor",
			Handler:    _ServerComms_ChangeSuccessor_Handler,
		},
		{
			MethodName: "changePredecessor",
			Handler:    _ServerComms_ChangePredecessor_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ServerComms_Get_Handler,
		},
		{
			MethodName: "put",
			Handler:    _ServerComms_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ServerComm.proto",
}
