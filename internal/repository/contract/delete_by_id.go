package contract

import "context"

type DeleteById[C context.Context] interface {
	DeleteById(C, int) error
}
