package user

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements UserRepository.
func (repo *UserRepositoryImpl) Save(ctx context.Context, user *domain.User) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsExist(repo.db, user.ID) {
		return errors.New("user already exists")
	}

	user.ID = repo.db[len(repo.db)].ID + 1
	repo.db[user.ID] = *user

	return nil
}
