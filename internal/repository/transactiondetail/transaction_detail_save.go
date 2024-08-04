package transactiondetail

import (
	"context"

	"gotik/internal/domain"
)

const create = `--
INSERT INTO
    transaction_details (time, status, total_payment)
VALUES
    ($1, $2, $3)
RETURNING
    id, time, status, total_payment
`

// Save implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) Save(ctx context.Context, td *domain.TransactionDetail) (transaction_detail domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	// if util.IsExist(repo.dbMap, td.ID) {
	// 	return domain.TransactionDetail{}, errors.New("transaction detail already exists")
	// }

	// td.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	// repo.dbMap[td.ID] = *td

	createStmt, err := repo.db.PrepareContext(ctx, create)
	if err != nil {
		return domain.TransactionDetail{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.TransactionDetail{}, err
	}

	row := tx.StmtContext(ctx, createStmt).QueryRowContext(ctx, &td.Time, &td.Status, &td.TotalPayment)
	err = row.Scan(
		&transaction_detail.ID,
		&transaction_detail.Time,
		&transaction_detail.Status,
		&transaction_detail.TotalPayment,
	)

	if err != nil {
		err = tx.Rollback()
		return domain.TransactionDetail{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.TransactionDetail{}, err
	}

	return transaction_detail, nil
}

const createTransactionDetailsEventsUsers = `--
INSERT INTO
    transaction_details_events_users (transaction_detail_id, event_id, user_id)
VALUES
    ($1, $2, $3)
`

// Save implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) SaveTransactionDetailsEventsUsers(ctx context.Context, tdID, eventID, userID int) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	createStmt, err := repo.db.PrepareContext(ctx, createTransactionDetailsEventsUsers)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.StmtContext(ctx, createStmt).ExecContext(ctx, tdID, eventID, userID)
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
