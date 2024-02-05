package inmemory

import (
	"context"
	"github.com/stretchr/testify/mock"
	"test/kit/command"
)

// MockCommandBus is a mock implementation of the CommandBus for testing.
type MockCommandBus struct {
	mock.Mock
}

// NewMockCommandBus initializes a new instance of MockCommandBus.
func NewMockCommandBus() *MockCommandBus {
	return &MockCommandBus{}
}

// Dispatch implements the CommandBus interface for the mock.
func (m *MockCommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

// Register implements the CommandBus interface for the mock.
func (m *MockCommandBus) Register(cmdType command.Type, handler command.Handler) {
	// Mock the Register method if needed.
}
