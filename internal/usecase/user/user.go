package user

import (
	"context"
	"gotik/internal/domain"
	"gotik/internal/repository/user"
)

type UserUsecase interface {
	user.UserRepository[context.Context, domain.User]
}

type UserUsecaseImpl struct {
	repo user.UserRepository[context.Context, domain.User]
}

func NewUserUsecase(repo user.UserRepository[context.Context, domain.User]) UserUsecase {
	return UserUsecaseImpl{
		repo: repo,
	}
}
