package data_structures

type SinglyLinkedListNode[T any] struct {
	val  T
	next *SinglyLinkedListNode[T]
}

func NewSinglyLinkedListNode[T any](val T) *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		val: val,
	}
}

func NewSentinelSinglyLinkedListNode[T any]() *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		val: *new(T),
	}
}

func NewSinglyLinkedList[T any](arr []T) *SinglyLinkedListNode[T] {
	headSentinel := NewSentinelSinglyLinkedListNode[T]()

	cur := headSentinel
	for _, e := range arr {
		cur.next = NewSinglyLinkedListNode[T](e)
		cur = cur.next
	}

	return headSentinel
}
