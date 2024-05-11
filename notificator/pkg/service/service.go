package service

import "context"

type SendRequest struct {
	Email string
	Name  string
	Amail string
	mail  string
}

type SendResponse struct {
	Data string
	Mata string
}

// NotificatorService describes the service.
type NotificatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmail(ctx context.Context, email string, content string) error

	Send(ctx context.Context, req SendRequest, content string) SendResponse
}

type basicNotificatorService struct{}

func (b *basicNotificatorService) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	// TODO implement the business logic of SendEmail
	return e0
}

// NewBasicNotificatorService returns a naive, stateless implementation of NotificatorService.
func NewBasicNotificatorService() NotificatorService {
	return &basicNotificatorService{}
}

// New returns a NotificatorService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificatorService {
	var svc NotificatorService = NewBasicNotificatorService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicNotificatorService) Send(ctx context.Context, req SendRequest, content string) (s0 SendResponse) {
	// TODO implement the business logic of Send
	return s0
}
