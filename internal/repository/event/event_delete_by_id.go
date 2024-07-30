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

	if !util.IsExist(repo.dbMap, id) {
		return errors.New("event not found")
	}

	delete(repo.dbMap, id)

	return nil
}
