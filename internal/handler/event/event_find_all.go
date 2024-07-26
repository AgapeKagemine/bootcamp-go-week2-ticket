package event

import (
	"context"
	"gotik/internal/domain"
)

// FindAll implements EventHandler.
func (h EventHandlerImpl) FindAll(ctx context.Context) (events []domain.Event, err error) {
	return h.uc.FindAll(ctx)
}
