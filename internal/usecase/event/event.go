package event

import (
	"gotik/internal/repository/event"
)

type EventUsecase interface {
	event.EventRepository
}

type EventUsecaseImpl struct {
	repo event.EventRepository
}

func NewEventUsecase(repo event.EventRepository) EventUsecase {
	return EventUsecaseImpl{
		repo: repo,
	}
}
