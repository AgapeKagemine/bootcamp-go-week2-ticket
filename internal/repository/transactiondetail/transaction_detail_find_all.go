package transactiondetail

import (
	"context"

	"gotik/internal/domain"
)

const findAll = `
SELECT
	id, time, status, total_payment
FROM
	transaction_details
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
		var td domain.TransactionDetail
		err = rows.Scan(
			&td.ID,
			&td.Time,
			&td.Status,
			&td.TotalPayment,
		)

		if err != nil {
			return nil, err
		}

		tds = append(tds, td)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tds, nil
}
