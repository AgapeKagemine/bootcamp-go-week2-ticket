package user

import (
	"context"

	"gotik/internal/domain"
)

// FindById implements UserUsecase.
func (uc *UserUsecaseImpl) FindById(ctx context.Context, id int) (user domain.User, err error) {
	return uc.repo.FindById(ctx, id)
}
