package event

import (
	"context"
	"gotik/internal/contract"
	"gotik/internal/domain"
	"sync"
)

type EventRepository[C context.Context, T domain.Event] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	contract.DeleteById[C]
}

type EventRepositoryImpl struct {
	db map[int]domain.Event
	*sync.Mutex
}

func NewEventRepository() EventRepository[context.Context, domain.Event] {
	return &EventRepositoryImpl{
		db:    make(map[int]domain.Event),
		Mutex: &sync.Mutex{},
	}
}
