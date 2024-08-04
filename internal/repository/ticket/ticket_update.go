package ticket

import (
	"context"
	"gotik/internal/domain"
)

const update = `
UPDATE
    tickets
SET
    stock = $2,
    type = $3,
    price = $4
WHERE 
    id = $1
RETURNING
    id, stock, type, price
`

func (repo *TicketRepositoryImpl) Update(ctx context.Context, t *domain.Ticket) (ticket domain.Ticket, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, Ticket.ID) {
	// 	return domain.Ticket{}, errors.New("Ticket not found")
	// }

	// repo.dbMap[Ticket.ID] = *Ticket

	updateStmt, err := repo.db.PrepareContext(ctx, update)
	if err != nil {
		return domain.Ticket{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Ticket{}, err
	}

	row := tx.StmtContext(ctx, updateStmt).QueryRowContext(ctx, t.ID, t.Stock, t.Type, t.Price)
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
