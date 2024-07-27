package event

import (
	"context"
	"gotik/internal/domain"
)

// Update implements EventUsecase.
func (uc *EventUsecaseImpl) Update(ctx context.Context, event *domain.Event) error {
	return uc.repo.Update(ctx, event)
}
