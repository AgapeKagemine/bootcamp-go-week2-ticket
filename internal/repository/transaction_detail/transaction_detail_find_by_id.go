package transaction_detail

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindById implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) FindById(ctx context.Context, id int) (td domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	if !util.IsExist(repo.db, id) {
		return domain.TransactionDetail{}, errors.New("transaction detail not found")
	}
	return repo.db[id], nil
}
