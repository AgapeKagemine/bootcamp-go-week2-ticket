package event

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/event"
	"gotik/internal/repository/transactiondetail"
	"gotik/internal/repository/user"
)

type EventUsecase interface {
	event.EventRepository[context.Context, domain.Event]
	BuyTicket
}

type EventUsecaseImpl struct {
	eventRepo event.EventRepository[context.Context, domain.Event]
	userRepo  user.UserRepository[context.Context, domain.User]
	tdRepo    transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

func NewEventUsecase(eventRepo event.EventRepository[context.Context, domain.Event], userRepo user.UserRepository[context.Context, domain.User], tdRepo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]) EventUsecase {
	return &EventUsecaseImpl{
		eventRepo: eventRepo,
		userRepo:  userRepo,
		tdRepo:    tdRepo,
	}
}
