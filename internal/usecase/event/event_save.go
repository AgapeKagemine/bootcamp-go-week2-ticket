package event

import (
	"context"
	"gotik/internal/domain"
)

// Save implements EventUsecase.
func (uc *EventUsecaseImpl) Save(ctx context.Context, event *domain.Event) error {
	return uc.repo.Save(ctx, event)
}
