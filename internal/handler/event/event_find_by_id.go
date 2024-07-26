package event

import (
	"context"
	"gotik/internal/domain"
)

// FindById implements EventHandler.
func (h EventHandlerImpl) FindById(ctx context.Context, id int) (event domain.Event, err error) {
	return h.uc.FindById(ctx, id)
}
