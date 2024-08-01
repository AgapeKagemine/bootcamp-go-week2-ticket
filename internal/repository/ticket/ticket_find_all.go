package ticket

import (
	"context"

	"gotik/internal/domain"
)

const findAll = `
SELECT
	id, stock, type, price
FROM
	tickets
`

// FindAll implements TicketRepository.
func (repo *TicketRepositoryImpl) FindAll(ctx context.Context) (tickets []domain.Ticket, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsEmpty(repo.dbMap) {
	// 	return nil, errors.New("no ticket found")
	// }

	// for _, ticket := range repo.dbMap {
	// 	tickets = append(tickets, ticket)
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
		var ticket domain.Ticket
		err = rows.Scan(
			&ticket.ID,
			&ticket.Stock,
			&ticket.Type,
			&ticket.Price,
		)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}
