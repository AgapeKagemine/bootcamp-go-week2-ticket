package user

import (
	"context"

	"gotik/internal/domain"
)

const create = `--
INSERT INTO
    users (username, phone, email, balance)
VALUES
    ($1, $2, $3, $4)
RETURNING
    id, username, phone, email, balance
`

// Save implements UserRepository.
func (repo *UserRepositoryImpl) Save(ctx context.Context, u *domain.User) (user domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsExist(repo.dbMap, user.ID) {
	// 	return errors.New("user already exists")
	// }

	// user.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	// repo.dbMap[user.ID] = *user

	createStmt, err := repo.db.PrepareContext(ctx, create)
	if err != nil {
		return domain.User{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.User{}, err
	}

	row := tx.StmtContext(ctx, createStmt).QueryRowContext(ctx, u.Username, u.Phone, u.Email, u.Balance)
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Phone,
		&user.Email,
		&user.Balance,
	)

	if err != nil {
		err = tx.Rollback()
		return domain.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
