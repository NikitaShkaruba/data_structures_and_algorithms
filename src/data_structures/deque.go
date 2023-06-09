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
	headSentinel, tailSentinel := NewDoublyLinkedList[T]([]T{})

	for _, n := range arr {
		newNode := NewDoublyLinkedListNode(n, tailSentinel.prev, tailSentinel)
		tailSentinel.prev.next, tailSentinel.prev = newNode, newNode
	}

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
	newNode := NewDoublyLinkedListNode(val, d.headSentinel, d.headSentinel.next)
	d.headSentinel.next, d.headSentinel.next.prev = newNode, newNode

	d.size++
}

// PushBack works in O(1) time, O(1) space
func (d *Deque[T]) PushBack(val T) {
	newNode := NewDoublyLinkedListNode(val, d.tailSentinel.prev, d.tailSentinel)
	d.tailSentinel.prev, d.tailSentinel.prev.next = newNode, newNode

	d.size++
}

// PopFront works in O(1) time, O(1) space
func (d *Deque[T]) PopFront() T {
	nodeToPop := d.headSentinel.next

	d.headSentinel.next, d.headSentinel.next.next.prev = d.headSentinel.next.next, d.headSentinel
	nodeToPop.next, nodeToPop.prev = nil, nil

	d.size--

	return nodeToPop.val
}

// PopBack works in O(1) time, O(1) space
func (d *Deque[T]) PopBack() T {
	nodeToPop := d.tailSentinel.prev

	d.tailSentinel.prev, d.tailSentinel.prev.prev.next = d.tailSentinel.prev.prev, d.tailSentinel
	nodeToPop.prev, nodeToPop.next = nil, nil

	d.size--

	return nodeToPop.val
}

// PeekFront works in O(1) time, O(1) space
func (d *Deque[T]) PeekFront() T {
	return d.headSentinel.next.val
}

// PeekBack works in O(1) time, O(1) space
func (d *Deque[T]) PeekBack() T {
	return d.tailSentinel.prev.val
}

// GetSize works in O(1) time, O(1) space
func (d *Deque[T]) GetSize() int {
	return d.size
}
