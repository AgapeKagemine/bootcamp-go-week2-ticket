package user

import (
	"context"

	"gotik/internal/domain"
)

// Update implements UserRepository.
func (repo *UserRepositoryImpl) Update(ctx context.Context, u *domain.User) (user domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, user.ID) {
	// 	return errors.New("user not found")
	// }

	// repo.dbMap[user.ID] = *user

	return domain.User{}, nil
}
