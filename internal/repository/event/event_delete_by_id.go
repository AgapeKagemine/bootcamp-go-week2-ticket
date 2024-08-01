package event

import (
	"context"
)

const deleteById = `
DELETE FROM 
	events
WHERE
	id = $1
`

// DeleteById implements EventRepository.
func (repo *EventRepositoryImpl) DeleteById(ctx context.Context, id int) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, id) {
	// 	return errors.New("event not found")
	// }

	// delete(repo.dbMap, id)

	deleteByIdStmt, err := repo.db.PrepareContext(ctx, deleteById)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.StmtContext(ctx, deleteByIdStmt).ExecContext(ctx, id)
	if err != nil {
		err = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
