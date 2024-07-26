package user

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/util"
)

// DeleteById implements UserRepository.
func (repo UserRepositoryImpl) DeleteById(ctx context.Context, id int) error {
	if !util.IsExist(repo.db, id) {
		return errors.New("user not found")
	}
	delete(repo.db, id)
	fmt.Println("user deleted successfully")
	return nil
}
