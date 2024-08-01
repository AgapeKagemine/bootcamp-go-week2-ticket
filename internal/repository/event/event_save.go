package event

import (
	"context"

	"gotik/internal/domain"
)

const create = `--
INSERT INTO
    events (id, name, date, description, location)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING
    id, name, date, description, location
`

// Save implements EventRepository.
func (repo *EventRepositoryImpl) Save(ctx context.Context, e *domain.Event) (event domain.Event, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsExist(repo.dbMap, event.ID) {
	// 	return domain.Event{}, errors.New("event already exists")
	// }

	// event.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	// repo.dbMap[event.ID] = *event

	createStmt, err := repo.db.PrepareContext(ctx, create)
	if err != nil {
		return domain.Event{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Event{}, err
	}

	row := tx.StmtContext(ctx, createStmt).QueryRowContext(ctx, e.ID, e.Name, e.Date, e.Description, e.Location)
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
