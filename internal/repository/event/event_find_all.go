package event

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindAll implements EventRepository.
func (repo EventRepositoryImpl) FindAll(ctx context.Context) (events []domain.Event, err error) {
	if util.IsEmpty(repo.db) {
		return nil, errors.New("no event found")
	}
	for _, event := range repo.db {
		events = append(events, event)
	}
	return events, nil
}
