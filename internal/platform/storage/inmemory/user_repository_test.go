package inmemory

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	domain "test/internal"
	"testing"
)

func Test_should_throw_error_if_user_not_found(t *testing.T) {
	repo := NewUserRepository()
	userID, _ := domain.NewUserID("c5a156c9-7878-47b1-bbaf-55f81204e65a")

	_, err := repo.FindOneById(userID)

	assert.Equal(t, domain.ErrUserNotFound, err)

}

func Test_should_return_user_by_id(t *testing.T) {
	repo := NewUserRepository()
	userVO, _ := domain.NewUser("c5a156c9-7878-47b1-bbaf-55f81204e65b", "best_player", 500)
	_ = repo.Save(userVO)

	user, _ := repo.FindOneById(userVO.ID())
	assert.Equal(t, "c5a156c9-7878-47b1-bbaf-55f81204e65b", user.ID().String())
	assert.Equal(t, "best_player", user.Username())
}

func Test_should_return_all_users(t *testing.T) {
	repo := NewUserRepository()
	users := repo.FindAll()
	assert.Equal(t, 10, len(users))
}

func Test_should_update_total_score_when_user_exists(t *testing.T) {
	repo := NewUserRepository()
	id, _ := domain.NewUserID("84aef8bc-69da-4753-83b0-1c48905c03a6")
	err := repo.UpdateTotalScore(id, 4500)
	require.NoError(t, err)

	user, _ := repo.FindOneById(id)
	assert.Equal(t, 4500, user.TotalScore())

}
