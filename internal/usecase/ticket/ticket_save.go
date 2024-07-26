package ticket

import (
	"context"
	"gotik/internal/domain"
)

// Save implements TicketUsecase.
func (uc TicketUsecaseImpl) Save(ctx context.Context, ticket *domain.Ticket) error {
	return uc.repo.Save(ctx, ticket)
}
