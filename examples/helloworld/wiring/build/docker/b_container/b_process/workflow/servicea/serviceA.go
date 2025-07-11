// Package servicea implements a simple Service called ServiceA that calls ServiceB
package servicea

import (
	"context"
	"fmt"
	"time"

	"github.com/blueprint-uservices/tutorial/examples/helloworld/workflow/serviceb"
)

// ServiceA provides the world-facing interface for service A
type ServiceA interface {
	Hello(ctx context.Context) error
	World(ctx context.Context) error
}

type ServiceAImpl struct {
	serviceB serviceb.ServiceB
}

func NewServiceA(ctx context.Context, serviceB serviceb.ServiceB) (*ServiceAImpl, error) {
	return &ServiceAImpl{serviceB}, nil
}

func (s *ServiceAImpl) Hello(ctx context.Context) error {
	fmt.Println("Hello Was Called We Sleep Here")
	time.Sleep(1 * time.Second)
	return nil
}

func (s *ServiceAImpl) World(ctx context.Context) error {
	return nil
}
