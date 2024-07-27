package ticket

import (
	"context"
	"gotik/internal/domain"
)

// FindAll implements TicketUsecase.
func (uc *TicketUsecaseImpl) FindAll(ctx context.Context) (tickets []domain.Ticket, err error) {
	return uc.repo.FindAll(ctx)
}
