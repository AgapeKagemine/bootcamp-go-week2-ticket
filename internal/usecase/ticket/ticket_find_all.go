package ticket

import (
	"context"

	"gotik/internal/domain"
)

// FindAll implements TicketUsecase.
func (uc *TicketUsecaseImpl) FindAll(ctx context.Context) (tickets []domain.Ticket, err error) {
	tickets, err = uc.repo.FindAll(ctx)
	if err != nil {
		return make([]domain.Ticket, 0), err
	}

	return tickets, nil
}
