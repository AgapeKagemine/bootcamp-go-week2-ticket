package event

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Update implements EventRepository.
func (repo *EventRepositoryImpl) Update(ctx context.Context, event *domain.Event) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if !util.IsExist(repo.db, event.ID) {
		return errors.New("event not found")
	}

	repo.db[event.ID] = *event

	return nil
}
