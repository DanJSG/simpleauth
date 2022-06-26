package collections

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}
