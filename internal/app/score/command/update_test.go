package command

import (
	"context"
	"github.com/stretchr/testify/require"
	domain "test/internal"
	"test/internal/platform/storage/inmemory"
	"testing"
)

func Test_should_perform_addition(t *testing.T) {

	userRepository := new(inmemory.MockUserRepository)
	sessionScoreCommandHandler := NewUpdateUserScoreCommandHandler(userRepository)
	user, _ := domain.NewUser("37a0f027-15e6-47cc-a5d2-64183281087e", "gio", 50)
	command := NewUpdateUserScoreCommand("37a0f027-15e6-47cc-a5d2-64183281087e", 50)
	userRepository.On("FindOneById", user.ID()).Return(user, nil)
	userRepository.On("UpdateTotalScore", user.ID(), uint(100)).Return(nil)
	//add other methods

	err := sessionScoreCommandHandler.Handle(context.Background(), command)
	require.NoError(t, err)

	userRepository.AssertExpectations(t)

}
