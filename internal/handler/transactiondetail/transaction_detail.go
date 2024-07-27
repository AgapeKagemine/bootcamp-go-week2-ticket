package transactiondetail

import (
	contract "gotik/internal/handler/contract/http"
	"gotik/internal/usecase/transactiondetail"
)

type TransactionDetailHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
}

type TransactionDetailHandlerImpl struct {
	uc transactiondetail.TransactionDetailUsecase
}

func NewTransactionDetailHandler(uc transactiondetail.TransactionDetailUsecase) TransactionDetailHandler {
	return &TransactionDetailHandlerImpl{
		uc: uc,
	}
}
