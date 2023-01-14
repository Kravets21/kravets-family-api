package services

import (
	"github.com/pkg/errors"
	"kravets-family-api/internal/users"
)

var (
	ErrRepositoryFindByID = func(err error, ID uint64) error {
		return errors.Wrapf(err, "findByID failed, ID = %d", ID)
	}
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) FindByID(id uint64) (users.User, error) {
	return users.User{
		ID:           id,
		Email:        "21kravets@gmail.com",
		PasswordHash: "qwerty",
		Username:     "vadim",
	}, nil

}
