package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsersService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "e0", e0)
	}()
	return l.next.Create(ctx, email)
}

func (l loggingMiddleware) Hello(ctx context.Context, name string) (s0 string) {
	defer func() {
		l.logger.Log("method", "Hello", "name", name, "s0", s0)
	}()
	return l.next.Hello(ctx, name)
}
func (l loggingMiddleware) Bye(ctx context.Context, name string) (s0 string) {
	defer func() {
		l.logger.Log("method", "Bye", "name", name, "s0", s0)
	}()
	return l.next.Bye(ctx, name)
}
