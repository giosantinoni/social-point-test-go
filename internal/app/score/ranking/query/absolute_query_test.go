package query

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	domain "test/internal"
	"test/internal/platform/storage/inmemory"
	"testing"
)

func Test_should_calculate_top_ranking(t *testing.T) {

	userRepository := new(inmemory.MockUserRepository)

	absoluteQueryHandler := NewAbsoluteRankingQueryHandler(userRepository)
	user1, _ := domain.NewUser("5280e6b3-488d-4975-a492-de6c72dd07df", "user_1st", 10)
	user2, _ := domain.NewUser("bdd3ea46-6c6a-4c93-aab4-3dea38079766", "user_2nd", 20)
	user3, _ := domain.NewUser("45d36a43-b747-41fc-b7e1-ffadab29b196", "user_3rd", 30)
	user4, _ := domain.NewUser("c9a4fb49-e744-438a-b7a6-651220700365", "user_4th", 40)
	user5, _ := domain.NewUser("c9485310-10b5-475a-84ef-8a5fe13a7edc", "user_5th", 50)
	users := append(make([]*domain.User, 0), &user1, &user2, &user3, &user4, &user5)

	userRepository.On("FindAll").Return(users)

	query := NewAbsoluteRankingQuery(3)
	ranking, err := absoluteQueryHandler.Handle(context.Background(), query)
	require.NoError(t, err)
	castedRanking := ranking.(domain.Ranking)
	userRepository.AssertExpectations(t)
	assert.Equal(t, 3, len(castedRanking.Users()))
	assert.Equal(t, 50, castedRanking.Users()[0].TotalScore())
	assert.Equal(t, 40, castedRanking.Users()[1].TotalScore())
	assert.Equal(t, 30, castedRanking.Users()[2].TotalScore())

}
