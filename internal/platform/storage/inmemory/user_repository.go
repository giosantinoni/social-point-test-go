package inmemory

import (
	domain "test/internal"
)

type UserRepository struct {
	users []*domain.User
}

// Es una función que crea y devuelve una nueva instancia del repositorio de usuarios
func NewUserRepository() *UserRepository {
	demoUser, _ := domain.NewUser("1c88e82c-5214-4d2c-8973-fa4b7d16dcb9", "FerozFantasma34", 0)
	demoUser1, _ := domain.NewUser("21717ff0-3426-422f-996a-d517380c358a", "FerozGuerrero31", 100)
	demoUser2, _ := domain.NewUser("84aef8bc-69da-4753-83b0-1c48905c03a6", "RápidoRayo1", 200)
	demoUser3, _ := domain.NewUser("01f9a6f9-4473-48ef-af36-26397aea8429", "RápidoDragón95", 300)
	demoUser4, _ := domain.NewUser("c24fb265-2f39-4a21-b058-c8fc23a649b5", "ÁgilPantera30", 400)
	demoUser5, _ := domain.NewUser("5280e6b3-488d-4975-a492-de6c72dd07df", "ÁgilGuerrero16", 500)
	demoUser6, _ := domain.NewUser("bdd3ea46-6c6a-4c93-aab4-3dea38079766", "OscuroRayo79", 600)
	demoUser7, _ := domain.NewUser("45d36a43-b747-41fc-b7e1-ffadab29b196", "RápidoRayo86", 700)
	demoUser8, _ := domain.NewUser("c9a4fb49-e744-438a-b7a6-651220700365", "RápidoDragón33", 800)
	demoUser9, _ := domain.NewUser("c9485310-10b5-475a-84ef-8a5fe13a7edc", "ValienteFantasma27", 900)

	users := []*domain.User{
		&demoUser, &demoUser1, &demoUser2, &demoUser3, &demoUser4, &demoUser5, &demoUser6, &demoUser7, &demoUser8, &demoUser9,
	}

	return &UserRepository{
		users: users,
	}
}

func (r *UserRepository) FindOneById(id domain.UserID) (domain.User, error) {
	for _, user := range r.users {
		if user.ID().String() == id.String() {
			return *user, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}

func (r *UserRepository) UpdateTotalScore(id domain.UserID, score uint) error {
	for key, existentUser := range r.users {
		if existentUser.ID().String() == id.String() {
			r.users[key].UpdateScore(score)
			return nil
		}
	}
	return domain.ErrUserNotFound
}

func (r *UserRepository) Save(user domain.User) error {
	r.users = append(r.users, &user)
	return nil
}

func (r *UserRepository) FindAll() []*domain.User {
	return r.users
}
