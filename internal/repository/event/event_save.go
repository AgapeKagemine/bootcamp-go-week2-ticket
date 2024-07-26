package event

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// Save implements EventRepository.
func (repo *EventRepositoryImpl) Save(ctx context.Context, event *domain.Event) error {
	if util.IsExist(repo.db, event.ID) {
		return errors.New("event already exists")
	}
	event.ID = repo.db[len(repo.db)].ID + 1
	repo.db[event.ID] = *event
	fmt.Println("event saved successfully")
	return nil
}
