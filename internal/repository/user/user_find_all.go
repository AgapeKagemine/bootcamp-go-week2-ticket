package user

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindAll implements UserRepository.
func (repo UserRepositoryImpl) FindAll(ctx context.Context) (users []domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	if util.IsEmpty(repo.db) {
		return nil, errors.New("no user found")
	}
	for _, user := range repo.db {
		users = append(users, user)
	}
	return users, nil
}
