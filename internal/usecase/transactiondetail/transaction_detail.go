package transactiondetail

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/repository/transactiondetail"
	"gotik/internal/repository/user"
)

type TransactionDetailUsecase interface {
	transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
}

type TransactionDetailUsecaseImpl struct {
	tdRepo   transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail]
	userRepo user.UserRepository[context.Context, domain.User]
}

func NewTransactionDetailUsecase(tdRepo transactiondetail.TransactionDetailRepository[context.Context, domain.TransactionDetail], userRepo user.UserRepository[context.Context, domain.User]) TransactionDetailUsecase {
	return &TransactionDetailUsecaseImpl{
		tdRepo:   tdRepo,
		userRepo: userRepo,
	}
}
