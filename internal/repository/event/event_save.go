package event

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements EventRepository.
func (repo *EventRepositoryImpl) Save(ctx context.Context, event *domain.Event) (domain.Event, error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsExist(repo.dbMap, event.ID) {
		return domain.Event{}, errors.New("event already exists")
	}

	event.ID = repo.dbMap[len(repo.dbMap)].ID + 1
	repo.dbMap[event.ID] = *event

	return domain.Event{}, nil
}
