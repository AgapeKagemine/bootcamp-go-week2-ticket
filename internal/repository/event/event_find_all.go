package event

import (
	"context"
	"errors"

	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindAll implements EventRepository.
func (repo *EventRepositoryImpl) FindAll(ctx context.Context) (events []domain.Event, err error) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if util.IsEmpty(repo.dbMap) {
		return nil, errors.New("no event found")
	}

	for _, event := range repo.dbMap {
		events = append(events, event)
	}

	return events, nil
}
