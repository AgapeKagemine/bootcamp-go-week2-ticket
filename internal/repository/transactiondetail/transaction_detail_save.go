package transactiondetail

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) Save(ctx context.Context, td *domain.TransactionDetail) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsExist(repo.db, td.ID) {
		return errors.New("transaction detail already exists")
	}

	td.ID = repo.db[len(repo.db)].ID + 1
	repo.db[td.ID] = *td

	return nil
}
