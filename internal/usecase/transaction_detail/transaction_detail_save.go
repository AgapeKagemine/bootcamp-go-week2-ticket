package transaction_detail

import (
	"context"
	"gotik/internal/domain"
)

// Save implements TransactionDetailUsecase.
func (uc TransactionDetailUsecaseImpl) Save(ctx context.Context, td *domain.TransactionDetail) error {
	return uc.repo.Save(ctx, td)
}
