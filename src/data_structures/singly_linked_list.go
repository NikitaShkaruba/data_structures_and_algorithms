package data_structures

////////////////////// Singly Linked List //////////////////////
// You can find this data structure overview in docs/data_structures/linked_list.md

// SinglyLinkedListNode is a linked list node with a pointer to the next node
type SinglyLinkedListNode[T any] struct {
	Val  T
	Next *SinglyLinkedListNode[T]
}

// NewSinglyLinkedListNode works in O(1) time, O(1) space
func NewSinglyLinkedListNode[T any](val T, next *SinglyLinkedListNode[T]) *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		Val:  val,
		Next: next,
	}
}

// NewSentinelSinglyLinkedListNode works in O(1) time, O(1) space
func NewSentinelSinglyLinkedListNode[T any]() *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		Val: *new(T),
	}
}

// NewSinglyLinkedList works in O(n) time, O(n) space
func NewSinglyLinkedList[T any](arr []T) *SinglyLinkedListNode[T] {
	headSentinel := NewSentinelSinglyLinkedListNode[T]()

	cur := headSentinel
	for _, e := range arr {
		cur.Next = NewSinglyLinkedListNode[T](e, nil)
		cur = cur.Next
	}

	return headSentinel
}
