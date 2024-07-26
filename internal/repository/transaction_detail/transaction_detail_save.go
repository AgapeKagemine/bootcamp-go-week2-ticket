package transaction_detail

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) Save(ctx context.Context, transaction_detail *domain.TransactionDetail) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	if util.IsExist(repo.db, transaction_detail.ID) {
		return errors.New("transaction detail already exists")
	}
	transaction_detail.ID = repo.db[len(repo.db)].ID + 1
	repo.db[transaction_detail.ID] = *transaction_detail
	fmt.Println("transaction detail saved successfully")
	return nil
}
