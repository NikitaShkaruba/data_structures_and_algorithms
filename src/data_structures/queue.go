package data_structures

////////////////////// Queue //////////////////////

// Queue is a data structure that implements LILO (last in, last out) interface
type Queue[T any] struct {
	deque *Deque[T]
}

// NewEmptyQueue works in O(1) time, O(1) space
func NewEmptyQueue[T any]() *Queue[T] {
	return NewQueueFromArray(make([]T, 0))
}

// NewQueueFromArray works in O(n) time, O(n) space
func NewQueueFromArray[T any](arr []T) *Queue[T] {
	return &Queue[T]{
		deque: NewDequeFromArray[T](arr),
	}
}

// Enqueue works in O(1) time, O(1) space
func (q *Queue[T]) Enqueue(val T) {
	q.deque.PushBack(val)
}

// Dequeue works in O(1) time, O(1) space
func (q *Queue[T]) Dequeue() T {
	return q.deque.PopFront()
}

// PeekFirst works in O(1) time, O(1) space
func (q *Queue[T]) PeekFirst() T {
	return q.deque.PeekFront()
}

// PeekLast works in O(1) time, O(1) space
func (q *Queue[T]) PeekLast() T {
	return q.deque.PeekBack()
}

// GetSize works in O(1) time, O(1) space
func (q *Queue[T]) GetSize() int {
	return q.deque.GetSize()
}
