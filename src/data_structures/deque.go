package data_structures

////////////////////// Deque //////////////////////

// Deque is a double ended queue that can push and pop from both ends for O(1) time, O(1) space
type Deque[T any] struct {
	headSentinel *DoublyLinkedListNode[T]
	tailSentinel *DoublyLinkedListNode[T]
	size         int
}

// NewDequeFromArray works in O(n) time, O(n) space
func NewDequeFromArray[T any](arr []T) *Deque[T] {
	headSentinel, tailSentinel := NewDoublyLinkedList(arr)

	return &Deque[T]{
		headSentinel: headSentinel,
		tailSentinel: tailSentinel,
		size:         len(arr),
	}
}

// NewEmptyDeque works in O(1) time, O(1) space
func NewEmptyDeque[T any]() *Deque[T] {
	return NewDequeFromArray(make([]T, 0))
}

// PushFront works in O(1) time, O(1) space
func (d *Deque[T]) PushFront(val T) {
	newNode := NewDoublyLinkedListNode(val, d.headSentinel, d.headSentinel.Next)
	d.headSentinel.Next, d.headSentinel.Next.Prev = newNode, newNode

	d.size++
}

// PushBack works in O(1) time, O(1) space
func (d *Deque[T]) PushBack(val T) {
	newNode := NewDoublyLinkedListNode(val, d.tailSentinel.Prev, d.tailSentinel)
	d.tailSentinel.Prev, d.tailSentinel.Prev.Next = newNode, newNode

	d.size++
}

// PopFront works in O(1) time, O(1) space
func (d *Deque[T]) PopFront() T {
	nodeToPop := d.headSentinel.Next

	d.headSentinel.Next, d.headSentinel.Next.Next.Prev = d.headSentinel.Next.Next, d.headSentinel
	nodeToPop.Next, nodeToPop.Prev = nil, nil

	d.size--

	return nodeToPop.Val
}

// PopBack works in O(1) time, O(1) space
func (d *Deque[T]) PopBack() T {
	nodeToPop := d.tailSentinel.Prev

	d.tailSentinel.Prev, d.tailSentinel.Prev.Prev.Next = d.tailSentinel.Prev.Prev, d.tailSentinel
	nodeToPop.Prev, nodeToPop.Next = nil, nil

	d.size--

	return nodeToPop.Val
}

// PeekFront works in O(1) time, O(1) space
func (d *Deque[T]) PeekFront() T {
	return d.headSentinel.Next.Val
}

// PeekBack works in O(1) time, O(1) space
func (d *Deque[T]) PeekBack() T {
	return d.tailSentinel.Prev.Val
}

// GetSize works in O(1) time, O(1) space
func (d *Deque[T]) GetSize() int {
	return d.size
}
