package event

import (
	"context"
	"gotik/internal/domain"
)

// FindAll implements EventUsecase.
func (uc *EventUsecaseImpl) FindAll(ctx context.Context) (events []domain.Event, err error) {
	return uc.repo.FindAll(ctx)
}
