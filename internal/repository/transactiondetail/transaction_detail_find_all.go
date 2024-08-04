package transactiondetail

import (
	"context"

	"gotik/internal/domain"
)

const findAllJoin = `
SELECT
	td.id, 
	td."time", 
	td.status, 
	td.total_payment, 
	u.id, 
	u.username, 
	u.phone, 
	u.email,
	u.balance,
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
	transaction_details td
JOIN
	transaction_details_events_users tdeu
	on
	tdeu.transaction_detail_id  = td.id
JOIN
	users u
	on
	tdeu.user_id = u.id
JOIN
	events e
	on
	tdeu.event_id = e.id
JOIN
	events_tickets et
	on
	e.id = et.event_id 
JOIN
	tickets t
	on
	et.ticket_id = t.id;
`

// FindAll implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) FindAll(ctx context.Context) (tds []domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsEmpty(repo.dbMap) {
	// 	return nil, errors.New("no transaction detail found")
	// }

	// for _, td := range repo.dbMap {
	// 	tds = append(tds, td)
	// }

	findAllStmt, err := repo.db.PrepareContext(ctx, findAllJoin)
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

	tdMap := make(map[int]*domain.TransactionDetail)

	for rows.Next() {
		td := &domain.TransactionDetail{}
		td.Event.Ticket = []domain.Ticket{}
		t := &domain.Ticket{}
		err = rows.Scan(
			&td.ID,
			&td.Time,
			&td.Status,
			&td.TotalPayment,
			&td.User.ID,
			&td.User.Username,
			&td.User.Phone,
			&td.User.Email,
			&td.User.Balance,
			&td.Event.ID,
			&td.Event.Name,
			&td.Event.Date,
			&td.Event.Description,
			&td.Event.Location,
			&t.ID,
			&t.Stock,
			&t.Type,
			&t.Price,
		)

		if err != nil {
			return nil, err
		}

		if _, exists := tdMap[td.ID]; !exists {
			tdMap[td.ID] = td
		}

		tdMap[td.ID].Event.Ticket = append(tdMap[td.ID].Event.Ticket, *t)

		if err != nil {
			return nil, err
		}
	}

	for _, td := range tdMap {
		tds = append(tds, *td)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tds, nil
}

const findAll = `
SELECT
	td.id, 
	td."time", 
	td.status, 
	td.total_payment
FROM
	transaction_details td
`

// FindAll implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) FindAllTrue(ctx context.Context) (tds []domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsEmpty(repo.dbMap) {
	// 	return nil, errors.New("no transaction detail found")
	// }

	// for _, td := range repo.dbMap {
	// 	tds = append(tds, td)
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
		td := &domain.TransactionDetail{}
		err = rows.Scan(
			&td.ID,
			&td.Time,
			&td.Status,
			&td.TotalPayment,
		)

		if err != nil {
			return nil, err
		}

		tds = append(tds, *td)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tds, nil
}
