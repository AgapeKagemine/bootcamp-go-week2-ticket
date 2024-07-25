package contract

type FindAll[T any] interface {
	FindAll() ([]T, error)
}
