package user

import (
	"gotik/internal/contract"
	"gotik/internal/domain"
)

type UserRepository interface {
	contract.FindAll[domain.User]
	contract.FindById[domain.User]
	contract.Save[domain.User]
	contract.Update[domain.User]
	contract.DeleteById[domain.User]
}

type UserRepositoryImpl struct {
	db map[int]domain.User
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{
		db: make(map[int]domain.User),
	}
}
