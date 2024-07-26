package contract

import "context"

type FindAll[C context.Context, T any] interface {
	FindAll(C) ([]T, error)
}
