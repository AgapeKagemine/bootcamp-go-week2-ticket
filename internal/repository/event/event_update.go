package event

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Update implements EventRepository.
func (repo *EventRepositoryImpl) Update(ctx context.Context, event *domain.Event) (domain.Event, error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if !util.IsExist(repo.dbMap, event.ID) {
		return domain.Event{}, errors.New("event not found")
	}

	repo.dbMap[event.ID] = *event

	return domain.Event{}, nil
}
