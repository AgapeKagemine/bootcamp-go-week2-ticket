package contract

import "context"

type Save[C context.Context, T any] interface {
	Save(C, *T) (T, error)
}
