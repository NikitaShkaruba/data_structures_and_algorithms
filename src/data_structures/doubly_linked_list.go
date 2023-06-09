package data_structures

////////////////////// Doubly Linked List //////////////////////

type DoublyLinkedListNode[T any] struct {
	val  T
	prev *DoublyLinkedListNode[T]
	next *DoublyLinkedListNode[T]
}

func NewDoublyLinkedListNode[T any](val T, prev *DoublyLinkedListNode[T], next *DoublyLinkedListNode[T]) *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		val:  val,
		prev: prev,
		next: next,
	}
}

func NewSentinelDoublyLinkedListNode[T any]() *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		val: *new(T),
	}
}

func NewDoublyLinkedList[T any](arr []T) (*DoublyLinkedListNode[T], *DoublyLinkedListNode[T]) {
	headSentinel := NewSentinelDoublyLinkedListNode[T]()
	tailSentinel := NewSentinelDoublyLinkedListNode[T]()
	headSentinel.next, tailSentinel.prev = tailSentinel, headSentinel

	cur := headSentinel
	for _, e := range arr {
		newNode := NewDoublyLinkedListNode[T](e, cur, cur.next)
		cur.next, cur.next.prev = newNode, newNode

		cur = cur.next
	}

	return headSentinel, tailSentinel
}
