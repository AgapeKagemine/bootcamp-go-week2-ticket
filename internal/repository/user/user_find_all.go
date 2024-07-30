package user

import (
	"context"

	"gotik/internal/domain"
)

const findAll = `
SELECT
	id, username, phone, email, balance
FROM
	users
`

// FindAll implements UserRepository.
func (repo *UserRepositoryImpl) FindAll(ctx context.Context) (users []domain.User, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsEmpty(repo.dbMap) {
	// 	return nil, errors.New("no user found")
	// }

	// for _, user := range repo.dbMap {
	// 	users = append(users, user)
	// }

	findAllStmt, err := repo.db.PrepareContext(ctx, findAll)
	if err != nil {
		return nil, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rows, err := tx.StmtContext(ctx, findAllStmt).QueryContext(ctx, findAll)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var i domain.User

		err = rows.Scan(
			&i.ID,
			&i.Username,
			&i.Phone,
			&i.Email,
			&i.Balance,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
