package user

import (
	"context"
)

// DeleteById implements UserUsecase.
func (uc *UserUsecaseImpl) DeleteById(ctx context.Context, id int) error {
	return uc.repo.DeleteById(ctx, id)
}
