package user

import (
	"context"
	"gotik/internal/domain"
)

// FindAll implements UserUsecase.
func (uc UserUsecaseImpl) FindAll(ctx context.Context) (users []domain.User, err error) {
	return uc.repo.FindAll(ctx)
}
