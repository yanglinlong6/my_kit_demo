package endpoint

import (
	"context"
	service "notificator/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// SendEmailRequest collects the request parameters for the SendEmail method.
type SendEmailRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// SendEmailResponse collects the response parameters for the SendEmail method.
type SendEmailResponse struct {
	E0 error `json:"e0"`
}

// MakeSendEmailEndpoint returns an endpoint that invokes SendEmail on the service.
func MakeSendEmailEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailRequest)
		e0 := s.SendEmail(ctx, req.Email, req.Content)
		return SendEmailResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SendEmailResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	request := SendEmailRequest{
		Content: content,
		Email:   email,
	}
	response, err := e.SendEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailResponse).E0
}

// SendRequest collects the request parameters for the Send method.
type SendRequest struct {
	Req     service.SendRequest `json:"req"`
	Content string              `json:"content"`
}

// SendResponse collects the response parameters for the Send method.
type SendResponse struct {
	S0 service.SendResponse `json:"s0"`
}

// MakeSendEndpoint returns an endpoint that invokes Send on the service.
func MakeSendEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		s0 := s.Send(ctx, req.Req, req.Content)
		return SendResponse{S0: s0}, nil
	}
}

// Send implements Service. Primarily useful in a client.
func (e Endpoints) Send(ctx context.Context, req service.SendRequest, content string) (s0 service.SendResponse) {
	request := SendRequest{
		Content: content,
		Req:     req,
	}
	response, err := e.SendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendResponse).S0
}
