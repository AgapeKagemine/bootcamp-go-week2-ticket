package transaction_detail

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindAll implements TransactionDetailRepository.
func (repo TransactionDetailRepositoryImpl) FindAll(ctx context.Context) (transaction_details []domain.TransactionDetail, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	if util.IsEmpty(repo.db) {
		return nil, errors.New("no transaction detail found")
	}
	for _, transaction_detail := range repo.db {
		transaction_details = append(transaction_details, transaction_detail)
	}
	return transaction_details, nil
}
