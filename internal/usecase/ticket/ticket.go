package ticket

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/ticket"
	"gotik/internal/usecase/contract"
)

type TicketUsecase[C context.Context, T domain.Ticket] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
}

type TicketUsecaseImpl struct {
	repo ticket.TicketRepository[context.Context, domain.Ticket]
}

func NewTicketUsecase(repo ticket.TicketRepository[context.Context, domain.Ticket]) TicketUsecase[context.Context, domain.Ticket] {
	return &TicketUsecaseImpl{
		repo: repo,
	}
}
