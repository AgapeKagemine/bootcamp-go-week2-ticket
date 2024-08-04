package event

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/event"
	"gotik/internal/repository/ticket"
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
	eventRepo  event.EventRepository[context.Context, domain.Event]
	userRepo   user.UserRepository[context.Context, domain.User]
	tdRepo     transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
	ticketRepo ticket.TicketRepository[context.Context, domain.Ticket]
}

func NewEventUsecase(
	eventRepo event.EventRepository[context.Context, domain.Event],
	userRepo user.UserRepository[context.Context, domain.User],
	tdRepo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail],
	ticketRepo ticket.TicketRepository[context.Context, domain.Ticket],
) EventUsecase[context.Context, domain.Event] {
	return &EventUsecaseImpl{
		eventRepo:  eventRepo,
		userRepo:   userRepo,
		tdRepo:     tdRepo,
		ticketRepo: ticketRepo,
	}
}
