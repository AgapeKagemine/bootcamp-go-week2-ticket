package transactiondetail

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/handler/contract"
	"gotik/internal/usecase/transactiondetail"
)

type TransactionDetailHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
}

type TransactionDetailHandlerImpl struct {
	uc transactiondetail.TransactionDetailUsecase[context.Context, domain.TransactionDetail]
}

func NewTransactionDetailHandler(uc transactiondetail.TransactionDetailUsecase[context.Context, domain.TransactionDetail]) TransactionDetailHandler {
	return &TransactionDetailHandlerImpl{
		uc: uc,
	}
}
