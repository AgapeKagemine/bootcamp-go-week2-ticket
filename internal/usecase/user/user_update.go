package user

import (
	"context"
	"gotik/internal/domain"
)

// Update implements UserUsecase.
func (uc UserUsecaseImpl) Update(ctx context.Context, user *domain.User) error {
	return uc.repo.Update(ctx, user)
}
