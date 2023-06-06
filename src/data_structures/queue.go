package data_structures

type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (s *Queue[T]) Enqueue(value T) {
	s.arr = append(s.arr, value)
}

func (s *Queue[T]) Dequeue() T {
	val := s.arr[0]
	s.arr = s.arr[1:]
	return val
}

func (s *Queue[T]) PeekFirst() T {
	return s.arr[0]
}

func (s *Queue[T]) PeekLast() T {
	return s.arr[len(s.arr)-1]
}

func (s *Queue[T]) GetSize() int {
	return len(s.arr)
}
