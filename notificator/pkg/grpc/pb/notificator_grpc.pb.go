// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: notificator/pkg/grpc/pb/notificator.proto

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

const (
	Notificator_SendEmail_FullMethodName   = "/pb.Notificator/SendEmail"
	Notificator_Send_FullMethodName        = "/pb.Notificator/Send"
	Notificator_FindByEmail_FullMethodName = "/pb.Notificator/FindByEmail"
	Notificator_FindById_FullMethodName    = "/pb.Notificator/FindById"
)

// NotificatorClient is the client API for Notificator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificatorClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailReply, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error)
	FindByEmail(ctx context.Context, in *FindByEmailRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FindById(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type notificatorClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificatorClient(cc grpc.ClientConnInterface) NotificatorClient {
	return &notificatorClient{cc}
}

func (c *notificatorClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailReply, error) {
	out := new(SendEmailReply)
	err := c.cc.Invoke(ctx, Notificator_SendEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificatorClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := c.cc.Invoke(ctx, Notificator_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificatorClient) FindByEmail(ctx context.Context, in *FindByEmailRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, Notificator_FindByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificatorClient) FindById(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, Notificator_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificatorServer is the server API for Notificator service.
// All implementations must embed UnimplementedNotificatorServer
// for forward compatibility
type NotificatorServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailReply, error)
	Send(context.Context, *SendRequest) (*SendReply, error)
	FindByEmail(context.Context, *FindByEmailRequest) (*UserResponse, error)
	FindById(context.Context, *FindByIDRequest) (*UserResponse, error)
	mustEmbedUnimplementedNotificatorServer()
}

// UnimplementedNotificatorServer must be embedded to have forward compatible implementations.
type UnimplementedNotificatorServer struct {
}

func (UnimplementedNotificatorServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedNotificatorServer) Send(context.Context, *SendRequest) (*SendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedNotificatorServer) FindByEmail(context.Context, *FindByEmailRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
}
func (UnimplementedNotificatorServer) FindById(context.Context, *FindByIDRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedNotificatorServer) mustEmbedUnimplementedNotificatorServer() {}

// UnsafeNotificatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificatorServer will
// result in compilation errors.
type UnsafeNotificatorServer interface {
	mustEmbedUnimplementedNotificatorServer()
}

func RegisterNotificatorServer(s grpc.ServiceRegistrar, srv NotificatorServer) {
	s.RegisterService(&Notificator_ServiceDesc, srv)
}

func _Notificator_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificatorServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notificator_SendEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificatorServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notificator_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificatorServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notificator_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificatorServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notificator_FindByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificatorServer).FindByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notificator_FindByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificatorServer).FindByEmail(ctx, req.(*FindByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notificator_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificatorServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notificator_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificatorServer).FindById(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notificator_ServiceDesc is the grpc.ServiceDesc for Notificator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notificator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Notificator",
	HandlerType: (*NotificatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _Notificator_SendEmail_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Notificator_Send_Handler,
		},
		{
			MethodName: "FindByEmail",
			Handler:    _Notificator_FindByEmail_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _Notificator_FindById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notificator/pkg/grpc/pb/notificator.proto",
}
