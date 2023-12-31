// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "users/pkg/endpoint"
	pb "users/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	create grpc.Handler
	hello  grpc.Handler
	bye    grpc.Handler
	pb.UnimplementedUsersServer
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.UsersServer {
	return &grpcServer{
		bye:    makeByeHandler(endpoints, options["Bye"]),
		create: makeCreateHandler(endpoints, options["Create"]),
		hello:  makeHelloHandler(endpoints, options["Hello"]),
	}
}
