package event

import (
	"context"

	"gotik/internal/domain"
)

const findById = `
SELECT
	e.id,
	e."name",
	e."date",
	e.description,
	e."location",
	t.id,
	t.stock,
	t."type",
	t.price 
FROM
	events e
JOIN 
	events_tickets et
	on
	e.id =  et.event_id
JOIN
	tickets t
	on
	et.ticket_id = t.id
WHERE
	e.id = $1
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

	rows, err := tx.StmtContext(ctx, findByIdStmt).QueryContext(ctx, id)
	if err != nil {
		return domain.Event{}, err
	}

	event.Ticket = []domain.Ticket{}

	for rows.Next() {
		var ticket domain.Ticket
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Date,
			&event.Description,
			&event.Location,
			&ticket.ID,
			&ticket.Stock,
			&ticket.Type,
			&ticket.Price,
		)
		event.Ticket = append(event.Ticket, ticket)
	}

	if err != nil {
		tx.Rollback()
		return domain.Event{}, err
	}

	err = rows.Close()
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}
