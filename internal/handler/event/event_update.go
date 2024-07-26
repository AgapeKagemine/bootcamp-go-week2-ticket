package event

import (
	"context"
	"gotik/internal/domain"
)

// Update implements EventHandler.
func (h EventHandlerImpl) Update(ctx context.Context, event *domain.Event) error {
	return h.uc.Update(ctx, event)
}
