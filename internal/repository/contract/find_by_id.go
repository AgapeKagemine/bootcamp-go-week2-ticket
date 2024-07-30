package contract

import "context"

type FindById[C context.Context, T any] interface {
	FindById(C, int) (T, error)
}
