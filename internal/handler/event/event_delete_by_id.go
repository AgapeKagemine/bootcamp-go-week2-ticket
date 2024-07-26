package event

import (
	"context"
)

// DeleteById implements EventHandler.
func (h EventHandlerImpl) DeleteById(ctx context.Context, id int) error {
	return h.uc.DeleteById(ctx, id)
}
