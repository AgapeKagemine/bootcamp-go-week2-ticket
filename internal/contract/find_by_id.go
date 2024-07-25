package contract

import "context"

type FindById[T any] interface {
	FindById(context.Context, int) (T, error)
}
