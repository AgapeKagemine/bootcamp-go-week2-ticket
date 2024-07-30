package event

import (
	"context"
	"database/sql"
	"sync"

	"gotik/internal/domain"
	"gotik/internal/repository/contract"
)

type EventRepository[C context.Context, T domain.Event] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	contract.DeleteById[C]
}

type EventRepositoryImpl struct {
	db    *sql.DB
	dbMap map[int]domain.Event
	*sync.Mutex
}

func NewEventRepository(db *sql.DB) EventRepository[context.Context, domain.Event] {
	return &EventRepositoryImpl{
		db:    db,
		dbMap: make(map[int]domain.Event),
		Mutex: &sync.Mutex{},
	}
}
