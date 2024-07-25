package event

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// Update implements EventRepository.
func (repo EventRepositoryImpl) Update(ctx context.Context, event *domain.Event) error {
	if !util.IsExist(repo.db, event.ID) {
		return errors.New("event not found")
	}
	repo.db[event.ID] = *event
	fmt.Println("event updated successfully")
	return nil
}
