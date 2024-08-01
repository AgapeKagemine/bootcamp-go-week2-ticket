package ticket

import (
	"context"

	"gotik/internal/domain"
)

const findById = `
SELECT
	id, stock, type, price
FROM
	tickets
WHERE
	id = $1
`

// FindById implements TicketRepository.
func (repo *TicketRepositoryImpl) FindById(ctx context.Context, id int) (ticket domain.Ticket, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, id) {
	// 	return domain.Ticket{}, errors.New("ticket not found")
	// }

	findByIdStmt, err := repo.db.PrepareContext(ctx, findById)
	if err != nil {
		return domain.Ticket{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Ticket{}, err
	}

	row := tx.StmtContext(ctx, findByIdStmt).QueryRowContext(ctx, id)
	err = row.Scan(
		&ticket.ID,
		&ticket.Stock,
		&ticket.Type,
		&ticket.Price,
	)

	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil
}
