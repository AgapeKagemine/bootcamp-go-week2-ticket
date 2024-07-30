package transactiondetail

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements TransactionDetailRepository.
func (repo *TransactionDetailRepositoryImpl) Save(ctx context.Context, td *domain.TransactionDetail) (domain.TransactionDetail, error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsExist(repo.dbMap, td.ID) {
		return domain.TransactionDetail{}, errors.New("transaction detail already exists")
	}

	td.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	repo.dbMap[td.ID] = *td

	return domain.TransactionDetail{}, nil
}
