package contract

import "context"

type Update[T any] interface {
	Update(context.Context, *T) error
}
