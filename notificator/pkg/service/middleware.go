package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NotificatorService) NotificatorService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificatorService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificatorService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificatorService) NotificatorService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "email", email, "content", content, "e0", e0)
	}()
	return l.next.SendEmail(ctx, email, content)
}

func (l loggingMiddleware) Send(ctx context.Context, req SendRequest, content string) (s0 SendResponse) {
	defer func() {
		l.logger.Log("method", "Send", "req", req, "content", content, "s0", s0)
	}()
	return l.next.Send(ctx, req, content)
}
