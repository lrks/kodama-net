package hello

import (
	"context"
	"fmt"
)

type Service interface {
	Greet(ctx context.Context, name string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Greet(_ context.Context, name string) (string, error) {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("Hello, %s!", name), nil
}
