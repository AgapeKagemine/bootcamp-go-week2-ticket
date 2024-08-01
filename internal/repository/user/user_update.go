package user

import (
	"context"

	"gotik/internal/domain"
)

const update = `
UPDATE
    users
SET
    username = $2,
    phone = $3,
    email = $4,
    balance = $5
WHERE 
    id = $1
RETURNING
    id, username, phone, email, balance
`

// Update implements UserRepository.
func (repo *UserRepositoryImpl) Update(ctx context.Context, u *domain.User) (user domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, user.ID) {
	// 	return errors.New("user not found")
	// }

	// repo.dbMap[user.ID] = *user

	updateStmt, err := repo.db.PrepareContext(ctx, update)
	if err != nil {
		return domain.User{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.User{}, err
	}

	row := tx.StmtContext(ctx, updateStmt).QueryRowContext(ctx, u.Username, u.Phone, u.Email, u.Balance)
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
