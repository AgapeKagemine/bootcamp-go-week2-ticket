package transactiondetail

import (
	"context"

	"gotik/internal/domain"
)

// FindById implements TransactionDetailUsecase.
func (uc *TransactionDetailUsecaseImpl) FindById(ctx context.Context, id int) (td domain.TransactionDetail, err error) {
	return uc.tdRepo.FindById(ctx, id)
}
