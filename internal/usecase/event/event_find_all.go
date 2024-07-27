package event

import (
	"context"

	"gotik/internal/domain"
)

// FindAll implements EventUsecase.
func (uc *EventUsecaseImpl) FindAll(ctx context.Context) (events []domain.Event, err error) {
	events, err = uc.repo.FindAll(ctx)
	if err != nil {
		return make([]domain.Event, 0), err
	}

	return events, err
}
