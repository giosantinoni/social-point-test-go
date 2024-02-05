package inmemory

import (
	"github.com/stretchr/testify/mock"
	"test/internal"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindOneById(id domain.UserID) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) Save(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateTotalScore(id domain.UserID, score uint) error {
	args := m.Called(id, score)
	return args.Error(0)
}

func (m *MockUserRepository) FindAll() []*domain.User {
	args := m.Called()
	if arg := args.Get(0); arg != nil {
		if users, ok := arg.([]*domain.User); ok {
			return users
		}
	}
	return nil
}
