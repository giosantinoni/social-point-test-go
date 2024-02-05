package query

import (
	"context"
	"errors"
	"sort"
	domain "test/internal"
	"test/kit/query"
)

const AbsoluteRankingQueryType query.Type = "query.absolute.ranking"

type AbsoluteRankingQuery struct {
	top int
}

func NewAbsoluteRankingQuery(top int) AbsoluteRankingQuery {
	return AbsoluteRankingQuery{
		top: top,
	}
}

func (q AbsoluteRankingQuery) Type() query.Type {
	return AbsoluteRankingQueryType
}

type AbsoluteRankingQueryHandler struct {
	userRepository domain.UserRepository
}

func NewAbsoluteRankingQueryHandler(userRepository domain.UserRepository) AbsoluteRankingQueryHandler {
	return AbsoluteRankingQueryHandler{
		userRepository: userRepository,
	}
}

func (h AbsoluteRankingQueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {
	q, ok := query.(AbsoluteRankingQuery)
	if !ok {
		return 0, errors.New("unexpected query")
	}

	ranking := h.userRepository.FindAll()

	sort.Sort(domain.Users(ranking))

	if len(ranking) > q.top {
		ranking = ranking[:q.top]
	}

	return domain.NewRanking(ranking, 0), nil
}
