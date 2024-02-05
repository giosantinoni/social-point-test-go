package command

import (
	"context"
	"errors"
	domain "test/internal"
	"test/kit/command"
)

const UpdateUserScoreCommandType command.Type = "command.update.totalScore"

// UpdateUserScoreCommand is the command dispatched to create a new course.
type UpdateUserScoreCommand struct {
	userID       string
	sessionScore int
}

func NewUpdateUserScoreCommand(userID string, score int) UpdateUserScoreCommand {
	return UpdateUserScoreCommand{
		userID:       userID,
		sessionScore: score,
	}
}

func (c UpdateUserScoreCommand) Type() command.Type {
	return UpdateUserScoreCommandType
}

// UpdateUserScoreCommandHandler is the command handler
// responsible for creating courses.
type UpdateUserScoreCommandHandler struct {
	userRepository domain.UserRepository
}

// NewUpdateUserScoreCommandHandler initializes a new UpdateUserScoreCommandHandler.
func NewUpdateUserScoreCommandHandler(userRepository domain.UserRepository) UpdateUserScoreCommandHandler {
	return UpdateUserScoreCommandHandler{
		userRepository: userRepository,
	}
}

// Handle implements the command.Handler interface.
func (h UpdateUserScoreCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	scoreCmd, ok := cmd.(UpdateUserScoreCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	id, err := domain.NewUserID(scoreCmd.userID)
	if err != nil {
		return err
	}

	user, err := h.userRepository.FindOneById(id)
	if err != nil {
		return err
	}

	totalScore := user.TotalScore() + scoreCmd.sessionScore

	if totalScore < 0 {
		totalScore = 0
	}

	return h.userRepository.UpdateTotalScore(id, uint(totalScore))

}
