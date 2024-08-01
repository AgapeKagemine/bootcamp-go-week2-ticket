package transactiondetail

import (
	"context"

	"gotik/internal/domain"
)

const findById = `
SELECT
	id, time, status, total_payment
FROM
	transaction_details
WHERE
	id = $1
`

// FindById implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) FindById(ctx context.Context, id int) (td domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if !util.IsExist(repo.dbMap, id) {
	// 	return domain.TransactionDetail{}, errors.New("transaction detail not found")
	// }

	findByIdStmt, err := repo.db.PrepareContext(ctx, findById)
	if err != nil {
		return domain.TransactionDetail{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.TransactionDetail{}, err
	}

	row := tx.StmtContext(ctx, findByIdStmt).QueryRowContext(ctx, id)
	err = row.Scan(
		&td.ID,
		&td.Time,
		&td.Status,
		&td.TotalPayment,
	)

	if err != nil {
		return domain.TransactionDetail{}, err
	}

	return td, nil
}
