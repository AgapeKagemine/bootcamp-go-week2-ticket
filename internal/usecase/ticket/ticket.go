package ticket

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/ticket"
)

type TicketUsecase interface {
	ticket.TicketRepository[context.Context, domain.Ticket]
}

type TicketUsecaseImpl struct {
	repo ticket.TicketRepository[context.Context, domain.Ticket]
}

func NewTicketUsecase(repo ticket.TicketRepository[context.Context, domain.Ticket]) TicketUsecase {
	return &TicketUsecaseImpl{
		repo: repo,
	}
}
