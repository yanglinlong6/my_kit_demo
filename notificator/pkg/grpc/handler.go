package grpc

import (
	"context"
	"errors"
	endpoint "notificator/pkg/endpoint"
	pb "notificator/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
)

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
// TODO implement the decoder
func decodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Decoder is not impelemented")
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Encoder is not impelemented")
}
func (g *grpcServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	_, rep, err := g.sendEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailReply), nil
}

// makeSendHandler creates the handler logic
func makeSendHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEndpoint, decodeSendRequest, encodeSendResponse, options...)
}

// decodeSendResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Send request.
// TODO implement the decoder
func decodeSendRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Decoder is not impelemented")
}

// encodeSendResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notificator' Encoder is not impelemented")
}
func (g *grpcServer) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendReply, error) {
	_, rep, err := g.send.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendReply), nil
}
