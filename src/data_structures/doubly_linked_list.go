package data_structures

////////////////////// Doubly Linked List //////////////////////

// DoublyLinkedListNode is a linked list node with pointers to the previous and next nodes
type DoublyLinkedListNode[T any] struct {
	Val  T
	Prev *DoublyLinkedListNode[T]
	Next *DoublyLinkedListNode[T]
}

// NewDoublyLinkedListNode works in O(1) time, O(1) space
func NewDoublyLinkedListNode[T any](val T, prev *DoublyLinkedListNode[T], next *DoublyLinkedListNode[T]) *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		Val:  val,
		Prev: prev,
		Next: next,
	}
}

// NewSentinelDoublyLinkedListNode works in O(1) time, O(1) space
func NewSentinelDoublyLinkedListNode[T any]() *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		Val: *new(T),
	}
}

// NewDoublyLinkedList works in O(n) time, O(n) space
func NewDoublyLinkedList[T any](arr []T) (*DoublyLinkedListNode[T], *DoublyLinkedListNode[T]) {
	headSentinel := NewSentinelDoublyLinkedListNode[T]()
	tailSentinel := NewSentinelDoublyLinkedListNode[T]()
	headSentinel.Next, tailSentinel.Prev = tailSentinel, headSentinel

	for _, n := range arr {
		newNode := NewDoublyLinkedListNode(n, tailSentinel.Prev, tailSentinel)
		tailSentinel.Prev.Next, tailSentinel.Prev = newNode, newNode
	}

	return headSentinel, tailSentinel
}
