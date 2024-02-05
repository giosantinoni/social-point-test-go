package command

import (
	"context"
	"github.com/stretchr/testify/require"
	domain "test/internal"
	"test/internal/platform/storage/inmemory"
	"testing"
)

func Test_should_not_perform_addition(t *testing.T) {

	userRepository := new(inmemory.MockUserRepository)
	sessionScoreCommandHandler := NewNewUserScoreCommandHandler(userRepository)
	user, _ := domain.NewUser("37a0f027-15e6-47cc-a5d2-64183281087e", "gio", 50)
	command := NewNewUserScoreCommand("37a0f027-15e6-47cc-a5d2-64183281087e", 50)

	userRepository.On("UpdateTotalScore", user.ID(), uint(50)).Return(nil)

	err := sessionScoreCommandHandler.Handle(context.Background(), command)
	require.NoError(t, err)

	userRepository.AssertExpectations(t)

}
