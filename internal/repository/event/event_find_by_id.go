package event

import (
	"context"

	"gotik/internal/domain"
)

const findById = `
SELECT
	id, name, date, description, location
FROM
	events
WHERE
	id = $1
`

// FindById implements EventRepository.
func (repo *EventRepositoryImpl) FindById(ctx context.Context, id int) (event domain.Event, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, id) {
	// 	return domain.Event{}, errors.New("event not found")
	// }

	findByIdStmt, err := repo.db.PrepareContext(ctx, findById)
	if err != nil {
		return domain.Event{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Event{}, err
	}

	row := tx.StmtContext(ctx, findByIdStmt).QueryRowContext(ctx, id)

	err = row.Scan(
		&event.ID,
		&event.Name,
		&event.Date,
		&event.Description,
	)

	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}
