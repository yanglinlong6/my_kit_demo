package service

import (
	"context"
	"fmt"
)

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, email string) error
	Hello(ctx context.Context, name string) string
	Bye(ctx context.Context, name string) string
}

type basicUsersService struct{}

func (b *basicUsersService) Create(ctx context.Context, email string) (e0 error) {
	// TODO implement the business logic of Create
	return e0
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUsersService) Hello(ctx context.Context, name string) (s0 string) {
	// TODO implement the business logic of Hello
	name = name + "ceshi"
	fmt.Println("name===" + name)
	return name + "Hello resp"
}
func (b *basicUsersService) Bye(ctx context.Context, name string) (s0 string) {
	// TODO implement the business logic of Bye
	return name + "Bye resp"
}
