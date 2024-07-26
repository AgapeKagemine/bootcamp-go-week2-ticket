package user

import (
	"context"
	"gotik/internal/domain"
)

// Save implements UserHandler.
func (h UserHandlerImpl) Save(ctx context.Context, user *domain.User) error {
	return h.uc.Save(ctx, user)
}
