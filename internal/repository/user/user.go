package user

import (
	"context"
	"sync"

	"gotik/internal/contract"
	"gotik/internal/domain"
)

type UserRepository[C context.Context, T domain.User] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	contract.DeleteById[C]
}

type UserRepositoryImpl struct {
	db map[int]domain.User
	*sync.Mutex
}

func NewUserRepository() UserRepository[context.Context, domain.User] {
	return &UserRepositoryImpl{
		db:    make(map[int]domain.User),
		Mutex: &sync.Mutex{},
	}
}

func (repo *UserRepositoryImpl) Get() UserRepositoryImpl {
	return *repo
}
