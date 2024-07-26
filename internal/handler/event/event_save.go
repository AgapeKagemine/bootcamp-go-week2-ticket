package event

import (
	"context"
	"gotik/internal/domain"
)

// Save implements EventHandler.
func (h EventHandlerImpl) Save(ctx context.Context, event *domain.Event) error {
	return h.uc.Save(ctx, event)
}
