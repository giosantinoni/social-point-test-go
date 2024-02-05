package query

import (
	"context"
	"errors"
	"sort"
	domain "test/internal"
	"test/kit/query"
)

const RelativeRankingQueryType query.Type = "query.relative.ranking"

type RelativeRankingQuery struct {
	rankingPosition int
	rankingOffset   int
}

// NewScoreService returns the default Service interface implementation.
func NewRelativeRankingQuery(rankingPosition, rankingOffset int) RelativeRankingQuery {
	return RelativeRankingQuery{
		rankingPosition: rankingPosition - 1,
		rankingOffset:   rankingOffset,
	}
}

func (q RelativeRankingQuery) Type() query.Type {
	return RelativeRankingQueryType
}

type RelativeRankingQueryHandler struct {
	userRepository domain.UserRepository
}

func NewRelativeRankingQueryHandler(userRepository domain.UserRepository) RelativeRankingQueryHandler {
	return RelativeRankingQueryHandler{
		userRepository: userRepository,
	}
}

func (h RelativeRankingQueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {
	relativeQuery, ok := query.(RelativeRankingQuery)
	if !ok {
		return nil, errors.New("unexpected query")
	}

	users := h.userRepository.FindAll()

	sort.Sort(domain.Users(users))
	if relativeQuery.rankingPosition >= len(users) {
		return nil, domain.ErrInvalidRankingRelativePosition
	}

	maxKey := relativeQuery.rankingPosition + relativeQuery.rankingOffset + 1
	if maxKey > len(users) {
		maxKey = len(users)
	}
	afterOffsetElements := users[relativeQuery.rankingPosition:maxKey]

	minKey := relativeQuery.rankingPosition - relativeQuery.rankingOffset
	if minKey < 0 {
		minKey = 0
	}
	beforeOffsetElements := users[minKey:relativeQuery.rankingPosition]

	ranking := append(beforeOffsetElements, afterOffsetElements...)

	return domain.NewRanking(ranking, minKey), nil

}
