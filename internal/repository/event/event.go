package event

import (
	"gotik/internal/contract"
	"gotik/internal/domain"
)

type EventRepository interface {
	contract.FindAll[domain.Event]
	contract.FindById[domain.Event]
	contract.Save[domain.Event]
	contract.Update[domain.Event]
	contract.DeleteById[domain.Event]
}

type EventRepositoryImpl struct {
	db map[int]domain.Event
}

func NewEventRepository() EventRepository {
	return EventRepositoryImpl{
		db: make(map[int]domain.Event),
	}
}
