package event

import (
	"context"

	"gotik/internal/domain"
)

const update = `
UPDATE
    events
SET
    name = $2,
    date = $3,
    description = $4,
    location = $5
WHERE 
    id = $1
RETURNING
    id, name, date, description, location
`

// Update implements EventRepository.
func (repo *EventRepositoryImpl) Update(ctx context.Context, e *domain.Event) (event domain.Event, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, event.ID) {
	// 	return domain.Event{}, errors.New("event not found")
	// }

	// repo.dbMap[event.ID] = *event

	updateStmt, err := repo.db.PrepareContext(ctx, update)
	if err != nil {
		return domain.Event{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Event{}, err
	}

	row := tx.StmtContext(ctx, updateStmt).QueryRowContext(ctx, e.ID, e.Name, e.Description, e.Location)
	err = row.Scan(
		&event.ID,
		&event.Name,
		&event.Date,
		&event.Description,
		&event.Location,
	)

	if err != nil {
		err = tx.Rollback()
		return domain.Event{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}
