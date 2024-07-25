package contract

type Save[T any] interface {
	Save(*T) error
}
