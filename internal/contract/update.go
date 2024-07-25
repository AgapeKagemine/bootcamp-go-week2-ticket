package contract

type Update[T any] interface {
	Update(*T) error
}
