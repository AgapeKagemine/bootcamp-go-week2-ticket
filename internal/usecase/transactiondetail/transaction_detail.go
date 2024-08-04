package transactiondetail

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/transactiondetail"
	"gotik/internal/usecase/contract"
)

type TransactionDetailUsecase[C context.Context, T domain.TransactionDetail] interface {
	contract.FindAll[C, T]
	contract.FindById[C, T]
	contract.Save[C, T]
}

type TransactionDetailUsecaseImpl struct {
	tdRepo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

func NewTransactionDetailUsecase(
	tdRepo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail],
) TransactionDetailUsecase[context.Context, domain.TransactionDetail] {
	return &TransactionDetailUsecaseImpl{
		tdRepo: tdRepo,
	}
}
