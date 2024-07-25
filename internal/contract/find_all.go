package contract

import "context"

type FindAll[T any] interface {
	FindAll(context.Context) ([]T, error)
}
