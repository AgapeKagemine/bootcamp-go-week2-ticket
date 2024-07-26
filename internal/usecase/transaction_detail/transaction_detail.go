package transaction_detail

import (
	"context"
	"gotik/internal/domain"
	"gotik/internal/repository/transaction_detail"
)

type TransactionDetailUsecase interface {
	transaction_detail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

type TransactionDetailUsecaseImpl struct {
	repo transaction_detail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

func NewTransactionDetailUsecase(repo transaction_detail.TransactionDetailRepository[context.Context, domain.TransactionDetail]) TransactionDetailUsecase {
	return TransactionDetailUsecaseImpl{
		repo: repo,
	}
}
