package ticket

import (
	"context"

	"gotik/internal/domain"
)

const create = `--
INSERT INTO
    tickets (stock, type, price)
VALUES
    ($1, $2, $3)
RETURNING
    id, stock, type, price
`

// Save implements TicketRepository.
func (repo *TicketRepositoryImpl) Save(ctx context.Context, t *domain.Ticket) (ticket domain.Ticket, err error) {
	// if util.IsExist(repo.dbMap, ticket.ID) {
	// 	return domain.Ticket{}, errors.New("ticket already exists")
	// }

	// ticket.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	// repo.dbMap[ticket.ID] = *ticket

	createStmt, err := repo.db.PrepareContext(ctx, create)
	if err != nil {
		return domain.Ticket{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Ticket{}, err
	}

	row := tx.StmtContext(ctx, createStmt).QueryRowContext(ctx, t.Stock, t.Type, t.Price)
	err = row.Scan(
		&ticket.ID,
		&ticket.Stock,
		&ticket.Type,
		&ticket.Price,
	)

	if err != nil {
		err = tx.Rollback()
		return domain.Ticket{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil
}

const createEventTicket = `--
INSERT INTO
    events_tickets (event_id, ticket_id)
VALUES
    ($1, $2)
`

func (repo *TicketRepositoryImpl) SaveEventTicket(ctx context.Context, ticketID int) error {
	createStmt, err := repo.db.PrepareContext(ctx, createEventTicket)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	eventID := ctx.Value(domain.Start("eventID")).(*int)

	row := tx.StmtContext(ctx, createStmt).QueryRowContext(ctx, *eventID, ticketID)

	err = row.Err()
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
