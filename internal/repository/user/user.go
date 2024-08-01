package user

import (
	"context"
	"database/sql"
	"sync"

	"gotik/internal/domain"
	"gotik/internal/repository/contract"
)

type UserRepository[C context.Context, T domain.User] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	contract.DeleteById[C]
}

type UserRepositoryImpl struct {
	db *sql.DB
	// dbMap map[int]domain.User
	*sync.Mutex
}

func NewUserRepository(db *sql.DB) UserRepository[context.Context, domain.User] {
	return &UserRepositoryImpl{
		db: db,
		// dbMap: make(map[int]domain.User),
		Mutex: &sync.Mutex{},
	}
}
