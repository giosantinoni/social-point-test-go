package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var ErrInvalidUserID = errors.New("invalid UUID for UserID")

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil
}

func (id UserID) String() string {
	return id.value
}

type UserRepository interface {
	FindOneById(id UserID) (User, error)
	UpdateTotalScore(id UserID, score uint) error
	Save(user User) error
	FindAll() []*User
}

var ErrUserNotFound = errors.New("The user was not found")

type User struct {
	id         UserID
	username   string
	totalScore uint
}

func NewUser(id string, name string, totalScore uint) (User, error) {
	idVO, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}
	return User{
		id:         idVO,
		username:   name,
		totalScore: totalScore,
	}, nil
}

func (u User) ID() UserID {
	return u.id
}

func (u User) Username() string {
	return u.username
}
func (u User) TotalScore() int {
	return int(u.totalScore)
}
func (u *User) UpdateScore(score uint) {
	u.totalScore = score
}

type Users []*User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	// Ordena de mayor a menor por totalScore
	return u[i].totalScore > u[j].totalScore
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
