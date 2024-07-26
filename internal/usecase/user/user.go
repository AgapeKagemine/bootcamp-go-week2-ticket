package user

import (
	"gotik/internal/repository/user"
)

type UserUsecase interface {
	user.UserRepository
}

type UserUsecaseImpl struct {
	repo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) UserUsecase {
	return UserUsecaseImpl{
		repo: repo,
	}
}
