package ticket

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindById implements TicketRepository.
func (repo TicketRepositoryImpl) FindById(ctx context.Context, id int) (ticket domain.Ticket, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	if !util.IsExist(repo.db, id) {
		return domain.Ticket{}, errors.New("ticket not found")
	}
	return repo.db[id], nil
}
