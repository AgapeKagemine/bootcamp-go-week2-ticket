package user

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/user"
	"gotik/internal/usecase/contract"
)

type UserUsecase[C context.Context, T domain.User] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	contract.DeleteById[C]
}

type UserUsecaseImpl struct {
	repo user.UserRepository[context.Context, domain.User]
}

func NewUserUsecase(repo user.UserRepository[context.Context, domain.User]) UserUsecase[context.Context, domain.User] {
	return &UserUsecaseImpl{
		repo: repo,
	}
}
