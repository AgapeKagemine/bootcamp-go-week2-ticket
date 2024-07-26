package contract

import "context"

type Update[C context.Context, T any] interface {
	Update(C, *T) error
}
