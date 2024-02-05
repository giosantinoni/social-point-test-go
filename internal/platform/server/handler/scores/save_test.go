package scores

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	appCommands "test/internal/app/score/command"
	"test/internal/platform/bus/inmemory"
	"testing"
)

var (
	emptyOperator = ScoreRequest{
		Score: 10, Operator: "",
	}
	minusOperator = ScoreRequest{
		Score: uint(10), Operator: "-",
	}
	equalOperator = ScoreRequest{
		Score: uint(10), Operator: "=",
	}
)

func Test_Save_Score_Handler(t *testing.T) {
	commandBus := new(inmemory.MockCommandBus)
	commandBus.On(
		"Dispatch",
		context.Background(),
		mock.Anything,
	).Return(nil)

	t.Run("given an invalid operator request returns 400", func(t *testing.T) {
		b, err := json.Marshal(emptyOperator)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/user/5280e6b3-488d-4975-a492-de6c72dd07df/score", bytes.NewBuffer(b))
		require.NoError(t, err)
		rr := httptest.NewRecorder()

		ScoreHandler(commandBus)(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func Test_map_update_score_command(t *testing.T) {

	commandBus := new(inmemory.MockCommandBus)
	r := mux.NewRouter()

	t.Run("given a request with MINUS operator dispatches correct command", func(t *testing.T) {

		command := appCommands.NewUpdateUserScoreCommand("5280e6b3-488d-4975-a492-de6c72dd07df", -10)
		commandBus.On(
			"Dispatch",
			mock.Anything,
			command,
		).Return(nil)

		r.HandleFunc("/user/{userId}/score", ScoreHandler(commandBus))
		b, err := json.Marshal(minusOperator)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/user/5280e6b3-488d-4975-a492-de6c72dd07df/score", bytes.NewBuffer(b))
		require.NoError(t, err)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
		commandBus.AssertExpectations(t)
	})

	t.Run("given a request with EQUAL operator dispatches correct command", func(t *testing.T) {

		command := appCommands.NewNewUserScoreCommand("5280e6b3-488d-4975-a492-de6c72dd07df", 10)
		commandBus.On(
			"Dispatch",
			mock.Anything,
			command,
		).Return(nil)

		r.HandleFunc("/user/{userId}/score", ScoreHandler(commandBus))
		b, err := json.Marshal(equalOperator)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/user/5280e6b3-488d-4975-a492-de6c72dd07df/score", bytes.NewBuffer(b))
		require.NoError(t, err)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
		commandBus.AssertExpectations(t)
	})
}
