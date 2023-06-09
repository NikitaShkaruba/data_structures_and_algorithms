package data_structures

////////////////////// Queue //////////////////////

type Queue[T any] struct {
	deque *Deque[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		deque: NewEmptyDeque[T](),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.deque.PushBack(val)
}

func (q *Queue[T]) Dequeue() T {
	return q.deque.PopFront()
}

func (q *Queue[T]) PeekFirst() T {
	return q.deque.PeekFront()
}

func (q *Queue[T]) PeekLast() T {
	return q.deque.PeekBack()
}

func (q *Queue[T]) GetSize() int {
	return q.deque.GetSize()
}
