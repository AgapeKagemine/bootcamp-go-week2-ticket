package transactiondetail

import (
	"context"

	"gotik/internal/domain"
)

// FindAll implements TransactionDetailUsecase.
func (uc *TransactionDetailUsecaseImpl) FindAll(ctx context.Context) (tds []domain.TransactionDetail, err error) {
	tds, err = uc.tdRepo.FindAll(ctx)
	if err != nil {
		return make([]domain.TransactionDetail, 0), err
	}

	return tds, nil
}
