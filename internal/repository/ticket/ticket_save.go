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
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

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

	row := tx.StmtContext(ctx, createStmt).QueryRowContext(ctx, t.ID, t.Stock, t.Type, t.Price)
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
