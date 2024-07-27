package user

import (
	"context"

	"gotik/internal/domain"
)

// FindAll implements UserUsecase.
func (uc *UserUsecaseImpl) FindAll(ctx context.Context) (users []domain.User, err error) {
	users, err = uc.repo.FindAll(ctx)
	if err != nil {
		return make([]domain.User, 0), err
	}

	return users, nil
}
