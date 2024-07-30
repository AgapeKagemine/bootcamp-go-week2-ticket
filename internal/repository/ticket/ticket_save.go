package ticket

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements TicketRepository.
func (repo *TicketRepositoryImpl) Save(ctx context.Context, ticket *domain.Ticket) (domain.Ticket, error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsExist(repo.dbMap, ticket.ID) {
		return domain.Ticket{}, errors.New("ticket already exists")
	}

	ticket.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	repo.dbMap[ticket.ID] = *ticket

	return domain.Ticket{}, nil
}
