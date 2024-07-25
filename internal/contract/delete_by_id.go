package contract

type DeleteById[T any] interface {
	DeleteById(int) (T, error)
}
