package user

import (
	"context"

	"gotik/internal/domain"
)

const findById = `
SELECT 
	id, username, phone, email, balance
FROM
	users
WHERE 
	id = $1
`

// FindById implements UserRepository.
func (repo *UserRepositoryImpl) FindById(ctx context.Context, id int) (user domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, id) {
	// 	return domain.User{}, errors.New("user not found")
	// }

	findByIdStmt, err := repo.db.PrepareContext(ctx, findById)
	if err != nil {
		return domain.User{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.User{}, err
	}

	row := tx.StmtContext(ctx, findByIdStmt).QueryRowContext(ctx, id)
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Phone,
		&user.Email,
		&user.Balance,
	)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
