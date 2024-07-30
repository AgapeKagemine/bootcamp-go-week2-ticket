package user

import (
	"context"

	"gotik/internal/domain"
)

// Save implements UserRepository.
func (repo *UserRepositoryImpl) Save(ctx context.Context, u *domain.User) (user domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsExist(repo.dbMap, user.ID) {
	// 	return errors.New("user already exists")
	// }

	// user.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	// repo.dbMap[user.ID] = *user

	return domain.User{}, nil
}
