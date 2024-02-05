package domain

import "errors"

type Ranking struct {
	users                []*User
	initialArrayPosition int
}

var ErrInvalidRankingRelativePosition = errors.New("invalid relative position for ranking. Not enough data")

func NewRanking(users []*User, position int) Ranking {
	return Ranking{
		users:                users,
		initialArrayPosition: position,
	}
}

func (r Ranking) Users() []*User {
	return r.users
}

func (r Ranking) InitialPosition() int {
	return r.initialArrayPosition + 1
}
