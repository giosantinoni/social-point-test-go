package scores

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"test/internal"
	appCommands "test/internal/app/score/command"
	"test/kit/command"
)

type ScoreRequest struct {
	Score    uint   `json:"score"`
	Operator string `json:"operator"`
}

var ErrInvalidOperator = errors.New("invalid Operator")

func ScoreHandler(commandBus command.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["userId"]
		var req ScoreRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd, err := mapToScoreCommand(req, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = commandBus.Dispatch(r.Context(), cmd)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidUserID):
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			case errors.Is(err, domain.ErrUserNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		_, err = fmt.Fprint(w, "Score submitted successfully")
		if err != nil {
			return
		}
	}
}

func mapToScoreCommand(saveScoreRequest ScoreRequest, userID string) (command.Command, error) {

	regex := regexp.MustCompile(`^[+\-=]$`)
	if !regex.MatchString(saveScoreRequest.Operator) {
		return nil,
			fmt.Errorf("%w: %s", ErrInvalidOperator, saveScoreRequest.Operator)
	}

	switch saveScoreRequest.Operator {
	case "-":
		return appCommands.NewUpdateUserScoreCommand(
			userID,
			int(saveScoreRequest.Score)*-1,
		), nil
	case "+":
		return appCommands.NewUpdateUserScoreCommand(
			userID,
			int(saveScoreRequest.Score),
		), nil
	default:
		return appCommands.NewNewUserScoreCommand(
			userID,
			saveScoreRequest.Score,
		), nil

	}
}
