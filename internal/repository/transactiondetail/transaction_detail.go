package transactiondetail

import (
	"context"
	"database/sql"
	"sync"

	"gotik/internal/domain"
	"gotik/internal/repository/contract"
)

type TransactionDetailRepository[C context.Context, T domain.TransactionDetail] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
	SaveTransactionDetailsEventsUsers(C, int, int, int) error
	FindAllTrue(C context.Context) ([]T, error)
}

type TransactionDetailRepositoryImpl struct {
	db *sql.DB
	// dbMap map[int]domain.TransactionDetail
	*sync.Mutex
}

func NewTransactionDetailRepository(db *sql.DB) TransactionDetailRepository[context.Context, domain.TransactionDetail] {
	return &TransactionDetailRepositoryImpl{
		db: db,
		// dbMap: make(map[int]domain.TransactionDetail),
		Mutex: &sync.Mutex{},
	}
}
