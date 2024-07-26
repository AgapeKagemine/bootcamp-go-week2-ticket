package user

import (
	"context"
	"gotik/internal/domain"
)

// FindById implements UserHandler.
func (h UserHandlerImpl) FindById(ctx context.Context, id int) (user domain.User, err error) {
	return h.uc.FindById(ctx, id)
}
