package transactiondetail

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/transactiondetail"
)

type TransactionDetailUsecase interface {
	transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

type TransactionDetailUsecaseImpl struct {
	repo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

func NewTransactionDetailUsecase(repo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]) TransactionDetailUsecase {
	return &TransactionDetailUsecaseImpl{
		repo: repo,
	}
}
