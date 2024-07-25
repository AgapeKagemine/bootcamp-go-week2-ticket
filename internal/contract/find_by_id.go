package contract

type FindById[T any] interface {
	FindById(T) (T, error)
}
