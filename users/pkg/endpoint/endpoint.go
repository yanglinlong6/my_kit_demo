package endpoint

import (
	"context"
	service "users/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Email string `json:"email"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		e0 := s.Create(ctx, req.Email)
		return CreateResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, email string) (e0 error) {
	request := CreateRequest{Email: email}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).E0
}

// HelloRequest collects the request parameters for the Hello method.
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse collects the response parameters for the Hello method.
type HelloResponse struct {
	S0 string `json:"s0"`
}

// MakeHelloEndpoint returns an endpoint that invokes Hello on the service.
func MakeHelloEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		s0 := s.Hello(ctx, req.Name)
		return HelloResponse{S0: s0}, nil
	}
}

// ByeRequest collects the request parameters for the Bye method.
type ByeRequest struct {
	Name string `json:"name"`
}

// ByeResponse collects the response parameters for the Bye method.
type ByeResponse struct {
	S0 string `json:"s0"`
}

// MakeByeEndpoint returns an endpoint that invokes Bye on the service.
func MakeByeEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ByeRequest)
		s0 := s.Bye(ctx, req.Name)
		return ByeResponse{S0: s0}, nil
	}
}

// Hello implements Service. Primarily useful in a client.
func (e Endpoints) Hello(ctx context.Context, name string) (s0 string) {
	request := HelloRequest{Name: name}
	response, err := e.HelloEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HelloResponse).S0
}

// Bye implements Service. Primarily useful in a client.
func (e Endpoints) Bye(ctx context.Context, name string) (s0 string) {
	request := ByeRequest{Name: name}
	response, err := e.ByeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ByeResponse).S0
}
