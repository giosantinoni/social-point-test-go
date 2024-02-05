package inmemory

import (
	"context"
	"test/kit/query"

	"github.com/stretchr/testify/mock"
)

// MockQueryBus is a mock implementation of the QueryBus interface.
type MockQueryBus struct {
	mock.Mock
}

// Dispatch implements the Dispatch method of the QueryBus interface for the mock.
func (m *MockQueryBus) Dispatch(ctx context.Context, q query.Query) (interface{}, error) {
	args := m.Called(ctx, q)
	return args.Get(0), args.Error(1)
}

// Register implements the Register method of the QueryBus interface for the mock.
func (m *MockQueryBus) Register(qType query.Type, handler query.Handler) {
	// You can add custom behavior here if needed
}
