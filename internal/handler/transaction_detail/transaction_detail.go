package transaction_detail

import (
	contract "gotik/internal/handler/contract/http"
	"gotik/internal/usecase/transaction_detail"
)

type TransactionDetailHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
}

type TransactionDetailHandlerImpl struct {
	uc transaction_detail.TransactionDetailUsecase
}

func NewTransactionDetailHandler(uc transaction_detail.TransactionDetailUsecase) TransactionDetailHandler {
	return &TransactionDetailHandlerImpl{
		uc: uc,
	}
}
