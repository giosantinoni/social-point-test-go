package scores

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	domain "test/internal"
	appQueries "test/internal/app/score/ranking/query"
	"test/kit/query"
)

var (
	topRegex = regexp.MustCompile(`^[Tt]op[1-9]\d*$`)
	atRegex  = regexp.MustCompile(`^At[1-9]\d*/[1-9]$`)
)

type RankingPresenter struct {
	Name     string
	Score    int
	Position int
}

func NewRankingPresenter(ranking domain.Ranking) []RankingPresenter {
	rankingPresenter := make([]RankingPresenter, 0)

	position := ranking.InitialPosition()
	for _, user := range ranking.Users() {
		rankingPresenter = append(rankingPresenter, RankingPresenter{user.Username(), user.TotalScore(), position})
		position++
	}

	return rankingPresenter
}

func RankingHandler(queryBus query.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		queryValues := r.URL.Query()
		rankingType := queryValues.Get("type")

		cqrsQuery, err := mapToRankingQuery(rankingType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ranking, err := queryBus.Dispatch(r.Context(), cqrsQuery)

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

		w.Header().Set("Content-Type", "application/json")

		rankingCasted := ranking.(domain.Ranking)
		presenter := NewRankingPresenter(rankingCasted)

		err = json.NewEncoder(w).Encode(presenter)

	}
}

func mapToRankingQuery(param string) (query.Query, error) {

	if topRegex.MatchString(param) {
		top, _ := strconv.Atoi(param[3:])
		return appQueries.NewAbsoluteRankingQuery(top), nil
	}

	if atRegex.MatchString(param) {
		parts := strings.Split(strings.TrimPrefix(param, "At"), "/")
		rankingPosition, _ := strconv.Atoi(parts[0])
		rankingOffset, _ := strconv.Atoi(parts[1])
		return appQueries.NewRelativeRankingQuery(rankingPosition, rankingOffset), nil
	}

	return nil, errors.New("invalid ranking format")
}
