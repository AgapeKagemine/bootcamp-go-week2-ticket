package user

import (
	"context"
	"gotik/internal/domain"
)

// Update implements UserHandler.
func (h UserHandlerImpl) Update(ctx context.Context, user *domain.User) error {
	return h.uc.Update(ctx, user)
}
