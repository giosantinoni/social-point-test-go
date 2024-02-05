package scores

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	domain "test/internal"
	appQueries "test/internal/app/score/ranking/query"
	"test/internal/platform/bus/inmemory"
	"testing"
)

// TODO: add a series of ranking variants to test
var rankingTypes = [3]string{"", "", ""}

func Test_map_ranking_query(t *testing.T) {

	queryBus := new(inmemory.MockQueryBus)
	r := mux.NewRouter()

	t.Run("given a request with valid Absolute <TopX> ranking dispatches correct command", func(t *testing.T) {

		query := appQueries.NewAbsoluteRankingQuery(100)
		queryBus.On(
			"Dispatch",
			mock.Anything,
			query,
		).Return(domain.NewRanking(make([]*domain.User, 0), 0), nil)

		r.HandleFunc("/ranking", RankingHandler(queryBus))

		req := httptest.NewRequest(http.MethodPost, "/ranking?type=Top100", bytes.NewBufferString(""))
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		queryBus.AssertExpectations(t)
	})

	t.Run("given a request with valid Relative <AtX/Y> operator dispatches correct command", func(t *testing.T) {

		query := appQueries.NewRelativeRankingQuery(5, 2)
		queryBus.On(
			"Dispatch",
			mock.Anything,
			query,
		).Return(domain.NewRanking(make([]*domain.User, 0), 0), nil)

		r.HandleFunc("/ranking", RankingHandler(queryBus))

		req := httptest.NewRequest(http.MethodPost, "/ranking?type=At5/2", bytes.NewBufferString(""))
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		queryBus.AssertExpectations(t)
	})

	t.Run("given a request with invalid ranking type returns Bad Request 400", func(t *testing.T) {

		query := appQueries.NewRelativeRankingQuery(5, 2)
		queryBus.On(
			"Dispatch",
			mock.Anything,
			query,
		).Return(domain.NewRanking(make([]*domain.User, 0), 0), nil)

		r.HandleFunc("/ranking", RankingHandler(queryBus))

		req := httptest.NewRequest(http.MethodPost, "/ranking?type=invalid-ranking", bytes.NewBufferString(""))
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})
}
