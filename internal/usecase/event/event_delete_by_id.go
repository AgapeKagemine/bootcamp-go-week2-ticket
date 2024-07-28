package event

import (
	"context"
)

// DeleteById implements EventUsecase.
func (uc *EventUsecaseImpl) DeleteById(ctx context.Context, id int) error {
	return uc.eventRepo.DeleteById(ctx, id)
}
