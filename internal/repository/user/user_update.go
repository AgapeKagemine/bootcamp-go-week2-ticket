package user

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// Update implements UserRepository.
func (repo UserRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	if !util.IsExist(repo.db, user.ID) {
		return errors.New("user not found")
	}
	repo.db[user.ID] = *user
	fmt.Println("user updated successfully")
	return nil
}
