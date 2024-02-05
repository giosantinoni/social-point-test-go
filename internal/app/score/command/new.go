package command

import (
	"context"
	"errors"
	domain "test/internal"
	"test/kit/command"
)

const NewUserScoreCommandType command.Type = "command.command.totalScore"

// NewUserScoreCommand is the command dispatched to create a new course.
type NewUserScoreCommand struct {
	userID     string
	totalScore uint
}

// UpdateScoreCommand creates a new NewUserScoreCommand.
func NewNewUserScoreCommand(userID string, score uint) NewUserScoreCommand {
	return NewUserScoreCommand{
		userID:     userID,
		totalScore: score,
	}
}

func (c NewUserScoreCommand) Type() command.Type {
	return NewUserScoreCommandType
}

// NewUserScoreCommandHandler is the command handler
// responsible for creating courses.
type NewUserScoreCommandHandler struct {
	userRepository domain.UserRepository
}

// NewNewUserScoreCommandHandler initializes a new NewUserScoreCommandHandler.
func NewNewUserScoreCommandHandler(userRepository domain.UserRepository) NewUserScoreCommandHandler {
	return NewUserScoreCommandHandler{
		userRepository: userRepository,
	}
}

// Handle implements the command.Handler interface.
func (h NewUserScoreCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createScoreCmd, ok := cmd.(NewUserScoreCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	id, err := domain.NewUserID(createScoreCmd.userID)
	if err != nil {
		return err
	}

	return h.userRepository.UpdateTotalScore(id, createScoreCmd.totalScore)

}
