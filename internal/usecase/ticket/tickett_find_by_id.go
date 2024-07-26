package ticket

import (
	"context"
	"gotik/internal/domain"
)

// FindById implements TicketUsecase.
func (uc TicketUsecaseImpl) FindById(ctx context.Context, id int) (ticket domain.Ticket, err error) {
	return uc.repo.FindById(ctx, id)
}
