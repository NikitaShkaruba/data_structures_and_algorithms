package data_structures

////////////////////// Stack //////////////////////
// You can find this data structure overview in docs/data_structures/stack.md
// TODO: start using linked list instead of a memory leaky array

type Stack[T any] struct {
	arr []T
}

// NewEmptyStack works in O(1) time, O(1) space
func NewEmptyStack[T any]() *Stack[T] {
	return NewStackFromArray(make([]T, 0))
}

// NewStackFromArray works in O(n) time, O(n) space
func NewStackFromArray[T any](arr []T) *Stack[T] {
	return &Stack[T]{
		arr: arr,
	}
}

// Push works in O(1) time, O(1) space
func (s *Stack[T]) Push(value T) {
	s.arr = append(s.arr, value)
}

// Pop works in O(1) time, O(1) space
func (s *Stack[T]) Pop() T {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

// Peek works in O(1) time, O(1) space
func (s *Stack[T]) Peek() T {
	return s.arr[len(s.arr)-1]
}

// GetSize works in O(1) time, O(1) space
func (s *Stack[T]) GetSize() int {
	return len(s.arr)
}
