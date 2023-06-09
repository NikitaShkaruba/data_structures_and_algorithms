package data_structures

////////////////////// Stack //////////////////////

type Stack[T any] struct {
	arr []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.arr = append(s.arr, value)
}

func (s *Stack[T]) Pop() T {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

func (s *Stack[T]) Peek() T {
	return s.arr[len(s.arr)-1]
}

func (s *Stack[T]) GetSize() int {
	return len(s.arr)
}
