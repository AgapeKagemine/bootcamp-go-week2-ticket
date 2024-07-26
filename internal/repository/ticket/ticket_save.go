package ticket

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements TicketRepository.
func (repo *TicketRepositoryImpl) Save(ctx context.Context, ticket *domain.Ticket) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	if util.IsExist(repo.db, ticket.ID) {
		return errors.New("ticket already exists")
	}
	ticket.ID = repo.db[len(repo.db)].ID + 1
	repo.db[ticket.ID] = *ticket
	fmt.Println("ticket saved successfully")
	return nil
}
