package contract

import "context"

type DeleteById[T any] interface {
	DeleteById(context.Context, int) error
}
