package event

import (
	"context"

	"gotik/internal/domain"
)

const findAll = `
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
	et.ticket_id = t.id;
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

	eventMap := make(map[int]*domain.Event)

	for rows.Next() {
		event := &domain.Event{}
		event.Ticket = []domain.Ticket{}
		ticket := &domain.Ticket{}
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

		if err != nil {
			return nil, err
		}

		if _, exists := eventMap[event.ID]; !exists {
			eventMap[event.ID] = event
		}

		eventMap[event.ID].Ticket = append(eventMap[event.ID].Ticket, *ticket)
	}

	for _, e := range eventMap {
		events = append(events, *e)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
