package service

import (
	"bugs/pkg/utils"
	"context"
	"fmt"
	npb "notificator/pkg/grpc/pb"
	"users/pkg/grpc/pb"
)

// BugsService describes the service.
type BugsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, bug string) error
}

type basicBugsService struct {
	//logger *zap.Logger
	userService        pb.UsersClient
	notificatorService npb.NotificatorClient
}

func (b *basicBugsService) Create(ctx context.Context, bug string) (e0 error) {
	// TODO implement the business logic of Create
	fmt.Println("bug===" + bug)
	//b.logger.Info("bug===" + bug)
	hello, _ := b.userService.Hello(ctx, &pb.HelloRequest{})
	fmt.Println("hello===" + hello.String())
	return e0
}

// NewBasicBugsService returns a naive, stateless implementation of BugsService.
func NewBasicBugsService(userService pb.UsersClient, notificatorService npb.NotificatorClient) BugsService {
	return &basicBugsService{
		userService:        userService,
		notificatorService: notificatorService,
	}
}

// New returns a BugsService with all of the expected middleware wired in.
func New(middleware []Middleware) BugsService {
	var svc BugsService = NewBasicBugsService(utils.CallOtherGRPCService(), utils.CallOtherGRPCService01())
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
