package event

import (
	"context"
	"errors"

	"gotik/internal/util"
)

// DeleteById implements EventRepository.
func (repo *EventRepositoryImpl) DeleteById(ctx context.Context, id int) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if !util.IsExist(repo.db, id) {
		return errors.New("event not found")
	}

	delete(repo.db, id)

	return nil
}
