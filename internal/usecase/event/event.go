package event

import (
	"context"
	"gotik/internal/domain"
	"gotik/internal/repository/event"
)

type EventUsecase interface {
	event.EventRepository[context.Context, domain.Event]
}

type EventUsecaseImpl struct {
	repo event.EventRepository[context.Context, domain.Event]
}

func NewEventUsecase(repo event.EventRepository[context.Context, domain.Event]) EventUsecase {
	return &EventUsecaseImpl{
		repo: repo,
	}
}
