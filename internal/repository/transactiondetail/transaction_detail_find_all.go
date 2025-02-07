package transactiondetail

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindAll implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) FindAll(ctx context.Context) (tds []domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsEmpty(repo.db) {
		return nil, errors.New("no transaction detail found")
	}

	for _, td := range repo.db {
		tds = append(tds, td)
	}

	return tds, nil
}
