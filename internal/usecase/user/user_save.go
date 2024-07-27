package user

import (
	"context"

	"gotik/internal/domain"
)

// Save implements UserUsecase.
func (uc *UserUsecaseImpl) Save(ctx context.Context, user *domain.User) error {
	return uc.repo.Save(ctx, user)
}
