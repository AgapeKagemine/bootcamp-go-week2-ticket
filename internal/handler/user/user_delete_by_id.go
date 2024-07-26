package user

import (
	"context"
)

// DeleteById implements UserHanlder.
func (h UserHandlerImpl) DeleteById(ctx context.Context, id int) error {
	return h.uc.DeleteById(ctx, id)
}
