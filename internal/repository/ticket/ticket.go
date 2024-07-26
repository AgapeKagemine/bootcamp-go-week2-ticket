package ticket

import (
	"context"
	"gotik/internal/contract"
	"gotik/internal/domain"
)

type TicketRepository[C context.Context, T domain.Ticket] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
}

type TicketRepositoryImpl struct {
	db map[int]domain.Ticket
}

func NewTicketRepository() TicketRepository[context.Context, domain.Ticket] {
	return &TicketRepositoryImpl{
		db: make(map[int]domain.Ticket),
	}
}
