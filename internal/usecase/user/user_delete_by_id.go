package user

import (
	"context"
)

// DeleteById implements UserRepository.
func (uc UserUsecaseImpl) DeleteById(ctx context.Context, id int) error {
	return uc.repo.DeleteById(ctx, id)
}
