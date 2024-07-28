package event

import (
	"context"

	"gotik/internal/domain"
)

// FindById implements EventUsecase.
func (uc *EventUsecaseImpl) FindById(ctx context.Context, id int) (event domain.Event, err error) {
	return uc.eventRepo.FindById(ctx, id)
}
