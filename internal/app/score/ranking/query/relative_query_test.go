package query

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	domain "test/internal"
	"test/internal/platform/storage/inmemory"
	"testing"
)

var (
	user1, _ = domain.NewUser("5280e6b3-488d-4975-a492-de6c72dd07df", "user_1st", 100)
	user2, _ = domain.NewUser("bdd3ea46-6c6a-4c93-aab4-3dea38079766", "user_2nd", 200)
	user3, _ = domain.NewUser("45d36a43-b747-41fc-b7e1-ffadab29b196", "user_3rd", 300)
	user4, _ = domain.NewUser("c9a4fb49-e744-438a-b7a6-651220700365", "user_4th", 400)
	user5, _ = domain.NewUser("c9485310-10b5-475a-84ef-8a5fe13a7edc", "user_5th", 500)
	user6, _ = domain.NewUser("c9485310-10b5-475a-84ef-8a5fe13a7edc", "user_6th", 600)
	users    = append(make([]*domain.User, 0), &user1, &user2, &user3, &user4, &user5, &user6)
)

func Test_should_calculate_relative_top_ranking_with_middle_offset(t *testing.T) {

	userRepository := new(inmemory.MockUserRepository)

	relativeQueryHandler := NewRelativeRankingQueryHandler(userRepository)
	userRepository.On("FindAll").Return(users)

	query := NewRelativeRankingQuery(3, 2)
	ranking, err := relativeQueryHandler.Handle(context.Background(), query)
	require.NoError(t, err)
	castedRanking := ranking.(domain.Ranking)
	userRepository.AssertExpectations(t)
	assert.Equal(t, 5, len(castedRanking.Users()))
	assert.Equal(t, 600, castedRanking.Users()[0].TotalScore())
	assert.Equal(t, 500, castedRanking.Users()[1].TotalScore())
	assert.Equal(t, 400, castedRanking.Users()[2].TotalScore())
	assert.Equal(t, 300, castedRanking.Users()[3].TotalScore())
	assert.Equal(t, 200, castedRanking.Users()[4].TotalScore())

}

func Test_should_calculate_relative_top_ranking_with_final_offset(t *testing.T) {

	userRepository := new(inmemory.MockUserRepository)

	relativeQueryHandler := NewRelativeRankingQueryHandler(userRepository)

	userRepository.On("FindAll").Return(users)

	query := NewRelativeRankingQuery(len(users), 2)
	ranking, err := relativeQueryHandler.Handle(context.Background(), query)
	require.NoError(t, err)
	castedRanking := ranking.(domain.Ranking)
	userRepository.AssertExpectations(t)
	require.Equal(t, 3, len(castedRanking.Users()))
	require.Equal(t, 300, castedRanking.Users()[0].TotalScore())
	require.Equal(t, 200, castedRanking.Users()[1].TotalScore())
	require.Equal(t, 100, castedRanking.Users()[2].TotalScore())

}

func Test_should_calculate_relative_top_ranking_with_init_offset(t *testing.T) {

	userRepository := new(inmemory.MockUserRepository)

	relativeQueryHandler := NewRelativeRankingQueryHandler(userRepository)

	userRepository.On("FindAll").Return(users)

	query := NewRelativeRankingQuery(1, 2)
	ranking, err := relativeQueryHandler.Handle(context.Background(), query)
	require.NoError(t, err)
	castedRanking := ranking.(domain.Ranking)
	userRepository.AssertExpectations(t)
	require.Equal(t, 3, len(castedRanking.Users()))
	require.Equal(t, 600, castedRanking.Users()[0].TotalScore())
	require.Equal(t, 500, castedRanking.Users()[1].TotalScore())
	require.Equal(t, 400, castedRanking.Users()[2].TotalScore())

}

func Test_should_throw_error_on_relative_position_greater_than_ranking(t *testing.T) {
	userRepository := new(inmemory.MockUserRepository)

	relativeQueryHandler := NewRelativeRankingQueryHandler(userRepository)

	userRepository.On("FindAll").Return(users)
	query := NewRelativeRankingQuery(7, 2)
	_, err := relativeQueryHandler.Handle(context.Background(), query)

	require.EqualError(t, domain.ErrInvalidRankingRelativePosition, err.Error())
	userRepository.AssertExpectations(t)

}
