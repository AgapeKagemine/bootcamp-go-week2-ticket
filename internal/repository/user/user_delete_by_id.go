package user

import (
	"context"
	"errors"

	"gotik/internal/util"
)

// DeleteById implements UserRepository.
func (repo *UserRepositoryImpl) DeleteById(ctx context.Context, id int) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if !util.IsExist(repo.db, id) {
		return errors.New("user not found")
	}

	delete(repo.db, id)

	return nil
}
