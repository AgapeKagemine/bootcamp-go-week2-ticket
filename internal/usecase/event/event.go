package event

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/event"
	"gotik/internal/repository/transactiondetail"
	"gotik/internal/repository/user"
	"gotik/internal/usecase/contract"
)

type EventUsecase[C context.Context, T domain.Event] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	contract.DeleteById[C]
	BuyTicket
}

type EventUsecaseImpl struct {
	eventRepo event.EventRepository[context.Context, domain.Event]
	userRepo  user.UserRepository[context.Context, domain.User]
	tdRepo    transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

func NewEventUsecase(eventRepo event.EventRepository[context.Context, domain.Event], userRepo user.UserRepository[context.Context, domain.User], tdRepo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]) EventUsecase[context.Context, domain.Event] {
	return &EventUsecaseImpl{
		eventRepo: eventRepo,
		userRepo:  userRepo,
		tdRepo:    tdRepo,
	}
}
