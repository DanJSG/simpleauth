package collections

import "fmt"

type Set[T comparable] interface {
	Iterable[T]
	Add(T)
	Remove(T)
	Contains(T) bool
	Size() int
}

type hashSet[T comparable] struct {
	data map[T]struct{}
}

type iterator[T comparable] struct {
	values *[]T
	length int
	i      int
}

var emptyStruct = struct{}{}

func HashSet[T comparable]() Set[T] {
	return &hashSet[T]{
		data: make(map[T]struct{}),
	}
}

func (s *hashSet[T]) Add(item T) {
	s.data[item] = emptyStruct
}

func (s *hashSet[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *hashSet[T]) Contains(item T) bool {
	_, found := s.data[item]
	return found
}

func (s *hashSet[T]) Size() int {
	return len(s.data)
}

func (s *hashSet[T]) String() string {
	values := getMapKeys[T](&s.data)
	return fmt.Sprintf("%#v", *values)
}

func (s *hashSet[T]) Iterator() Iterator[T] {
	values := getMapKeys[T](&s.data)
	return &iterator[T]{
		values: values,
		length: len(*values),
		i:      0,
	}
}

func getMapKeys[T comparable](m *map[T]struct{}) *[]T {
	i := 0
	keys := make([]T, len(*m))
	for key := range *m {
		keys[i] = key
		i++
	}
	return &keys
}

func (iter *iterator[T]) HasNext() bool {
	if iter.i < iter.length {
		return true
	} else {
		return false
	}
}

func (iter *iterator[T]) Next() T {
	value := (*iter.values)[iter.i]
	iter.i++
	return value
}
