package event

import (
	"context"
	"errors"
	"gotik/internal/domain"
	"gotik/internal/util"
)

// FindById implements EventRepository.
func (repo EventRepositoryImpl) FindById(ctx context.Context, id int) (event domain.Event, err error) {
	if !util.IsExist(repo.db, id) {
		return domain.Event{}, errors.New("event not found")
	}
	return repo.db[id], nil
}
