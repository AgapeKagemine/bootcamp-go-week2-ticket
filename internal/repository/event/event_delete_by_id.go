package event

import (
	"context"
	"errors"
	"fmt"
	"gotik/internal/util"
)

// DeleteById implements EventRepository.
func (repo *EventRepositoryImpl) DeleteById(ctx context.Context, id int) error {
	if !util.IsExist(repo.db, id) {
		return errors.New("event not found")
	}
	delete(repo.db, id)
	fmt.Println("event deleted successfully")
	return nil
}
