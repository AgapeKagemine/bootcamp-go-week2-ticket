package user

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements UserRepository.
func (repo UserRepositoryImpl) Save(ctx context.Context, user *domain.User) error {
	if util.IsExist(repo.db, user.ID) {
		return errors.New("user already exists")
	}
	user.ID = repo.db[len(repo.db)].ID + 1
	repo.db[user.ID] = *user
	fmt.Println("user saved successfully")
	return nil
}
