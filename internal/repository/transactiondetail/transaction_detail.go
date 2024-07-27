package transactiondetail

import (
	"context"
	"sync"

	"gotik/internal/contract"
	"gotik/internal/domain"
)

type TransactionDetailRepository[C context.Context, T domain.TransactionDetail] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
}

type TransactionDetailRepositoryImpl struct {
	db map[int]domain.TransactionDetail
	*sync.Mutex
}

func NewTransactionDetailRepository() TransactionDetailRepository[context.Context, domain.TransactionDetail] {
	return &TransactionDetailRepositoryImpl{
		db:    make(map[int]domain.TransactionDetail),
		Mutex: &sync.Mutex{},
	}
}
