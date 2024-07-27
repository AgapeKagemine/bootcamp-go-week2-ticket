package user

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindById implements UserRepository.
func (repo *UserRepositoryImpl) FindById(ctx context.Context, id int) (user domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if !util.IsExist(repo.db, id) {
		return domain.User{}, errors.New("user not found")
	}

	return repo.db[id], nil
}
