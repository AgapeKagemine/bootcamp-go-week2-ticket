package contract

import "context"

type Save[T any] interface {
	Save(context.Context, *T) error
}
