package ticket

import (
	"context"
	"database/sql"
	"sync"

	"gotik/internal/domain"
	"gotik/internal/repository/contract"
)

type TicketRepository[C context.Context, T domain.Ticket] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	contract.Update[C, T]
	SaveEventTicket(C, int) error
}

type TicketRepositoryImpl struct {
	db *sql.DB
	// dbMap map[int]domain.Ticket
	*sync.Mutex
}

func NewTicketRepository(db *sql.DB) TicketRepository[context.Context, domain.Ticket] {
	return &TicketRepositoryImpl{
		db: db,
		// dbMap: make(map[int]domain.Ticket),
		Mutex: &sync.Mutex{},
	}
}
