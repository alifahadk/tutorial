// Package servicea implements a simple Service called that calls ServiceB
package serviceb

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

// ServiceB provides the world-facing interface for service B
type ServiceB interface {
	Join(ctx context.Context) error
}

type ServiceBImpl struct {
	cache backend.Cache
}

func NewServiceB(ctx context.Context, cache backend.Cache) (*ServiceBImpl, error) {
	return &ServiceBImpl{cache}, nil
}

func (s *ServiceBImpl) Join(ctx context.Context) error {
	return nil
}
