package transaction_detail

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindById implements TransactionDetailRepository.
func (repo TransactionDetailRepositoryImpl) FindById(ctx context.Context, id int) (transaction_detail domain.TransactionDetail, err error) {
	if !util.IsExist(repo.db, id) {
		return domain.TransactionDetail{}, errors.New("transaction detail not found")
	}
	return repo.db[id], nil
}
