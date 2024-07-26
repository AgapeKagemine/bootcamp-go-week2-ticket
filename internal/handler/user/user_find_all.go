package user

import (
	"context"
	"gotik/internal/domain"
)

// FindAll implements UserHandler.
func (h UserHandlerImpl) FindAll(ctx context.Context) (users []domain.User, err error) {
	return h.uc.FindAll(ctx)
}
