package transaction_detail

import (
	"context"
	"gotik/internal/domain"
)

// FindAll implements TransactionDetailUsecase.
func (uc *TransactionDetailUsecaseImpl) FindAll(ctx context.Context) (tds []domain.TransactionDetail, err error) {
	return uc.repo.FindAll(ctx)
}
