package user

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Update implements UserRepository.
func (repo *UserRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if !util.IsExist(repo.db, user.ID) {
		return errors.New("user not found")
	}

	repo.db[user.ID] = *user

	return nil
}
