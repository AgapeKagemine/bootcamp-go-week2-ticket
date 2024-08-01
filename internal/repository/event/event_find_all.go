package event

import (
	"context"

	"gotik/internal/domain"
)

const findAll = `
SELECT
	id, name, date, description, location
FROM
	events
`

// FindAll implements EventRepository.
func (repo *EventRepositoryImpl) FindAll(ctx context.Context) (events []domain.Event, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsEmpty(repo.dbMap) {
	// 	return nil, errors.New("no event found")
	// }

	// for _, event := range repo.dbMap {
	// 	events = append(events, event)
	// }

	findAllStmt, err := repo.db.PrepareContext(ctx, findAll)
	if err != nil {
		return nil, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rows, err := tx.StmtContext(ctx, findAllStmt).QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event domain.Event
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Date,
			&event.Description,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
