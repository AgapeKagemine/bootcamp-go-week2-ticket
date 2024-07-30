package ticket

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindAll implements TicketRepository.
func (repo *TicketRepositoryImpl) FindAll(ctx context.Context) (tickets []domain.Ticket, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsEmpty(repo.dbMap) {
		return nil, errors.New("no ticket found")
	}

	for _, ticket := range repo.dbMap {
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
