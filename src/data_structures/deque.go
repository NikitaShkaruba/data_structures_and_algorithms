package data_structures

////////////////////// Deque //////////////////////

type Deque[T any] struct {
	headSentinel *DoublyLinkedListNode[T]
	tailSentinel *DoublyLinkedListNode[T]
	size         int
}

func NewDeque[T any]() *Deque[T] {
	headSentinel, tailSentinel := NewDoublyLinkedList[T]([]T{})

	return &Deque[T]{
		headSentinel: headSentinel,
		tailSentinel: tailSentinel,
	}
}

func (d *Deque[T]) PushFront(val T) {
	newNode := NewDoublyLinkedListNode(val, d.headSentinel, d.headSentinel.next)
	d.headSentinel.next, d.headSentinel.next.prev = newNode, newNode

	d.size++
}

func (d *Deque[T]) PushBack(val T) {
	newNode := NewDoublyLinkedListNode(val, d.tailSentinel.prev, d.tailSentinel)
	d.tailSentinel.prev, d.tailSentinel.prev.next = newNode, newNode

	d.size++
}

func (d *Deque[T]) PopFront() T {
	nodeToPop := d.headSentinel.next

	d.headSentinel.next, d.headSentinel.next.next.prev = d.headSentinel.next.next, d.headSentinel
	nodeToPop.next, nodeToPop.prev = nil, nil

	d.size--

	return nodeToPop.val
}

func (d *Deque[T]) PopBack() T {
	nodeToPop := d.tailSentinel.prev

	d.tailSentinel.prev, d.tailSentinel.prev.prev.next = d.tailSentinel.prev.prev, d.tailSentinel
	nodeToPop.prev, nodeToPop.next = nil, nil

	d.size--

	return nodeToPop.val
}

func (d *Deque[T]) PeekFront() T {
	return d.headSentinel.next.val
}

func (d *Deque[T]) PeekBack() T {
	return d.tailSentinel.prev.val
}

func (d *Deque[T]) GetSize() int {
	return d.size
}
