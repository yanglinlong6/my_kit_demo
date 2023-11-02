package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "users/pkg/endpoint"
	pb "users/pkg/grpc/pb"
)

// makeCreateHandler creates the handler logic
func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

// decodeCreateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Create request.
// TODO implement the decoder
func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	_, rep, err := g.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateReply), nil
}

// makeHelloHandler creates the handler logic
func makeHelloHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.HelloEndpoint, decodeHelloRequest, encodeHelloResponse, options...)
}

// decodeHelloResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Hello request.
// TODO implement the decoder
func decodeHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeHelloResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	_, rep, err := g.hello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloReply), nil
}

// makeByeHandler creates the handler logic
func makeByeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ByeEndpoint, decodeByeRequest, encodeByeResponse, options...)
}

// decodeByeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Bye request.
// TODO implement the decoder
func decodeByeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeByeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeByeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) Bye(ctx context.Context, req *pb.ByeRequest) (*pb.ByeReply, error) {
	_, rep, err := g.bye.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ByeReply), nil
}
