# Linked list

Can be used to store continuous data with slow access by index in `O(n)` time,
but very quick deletion and insertion from every index in `O(1)` time.

Linked list can have a pointer only to the next (singly linked list), or to the next and previous elements (doubly linked list)

Source code: [src/data_structures/singly_linked_list.go](../../src/data_structures/singly_linked_list.go), [src/data_structures/doubly_linked_list.go](../../src/data_structures/doubly_linked_list.go)

### Sentinels

Sentinels are helper nodes that stay at the beginning and ending of the linked list.
They don't store any data and are not counted as real linked list nodes.
They help to deal with edge cases where the first or the last element can be nil.
With sentinels, we’ll always have the first and the last, so we don’t deal with the nil problems.
It's a very good pattern.

```go
type LinkedList struct {
	headSentinel *data_structures.DoublyLinkedListNode[int]
	tailSentinel *data_structures.DoublyLinkedListNode[int]
	length       int
}

func NewLinkedList() *LinkedList {
	headSentinel := &data_structures.DoublyLinkedListNode[int]{}
	tailSentinel := &data_structures.DoublyLinkedListNode[int]{}

	headSentinel.Next = tailSentinel
	tailSentinel.Prev = headSentinel

	return &LinkedList{
		headSentinel: headSentinel,
		tailSentinel: tailSentinel,
	}
}

// Get works in O(n) time, O(1) space
func (l *LinkedList) Get(index int) int {
	if index < 0 || index >= l.length {
		return -1
	}

	// We don't need to check for nil, even if linked list's size is 0
	i := 0
	cur := l.headSentinel.Next
	for i != index {
		cur = cur.Next
		i++
	}

	return cur.Val
}

// AddAtHead works in O(1) time, O(1) space
func (l *LinkedList) AddAtHead(val int) {
	newNode := &data_structures.DoublyLinkedListNode[int]{
		Val:  val,
		Prev: l.headSentinel,
		Next: l.headSentinel.Next,
	}

	// We don't need to check for nil, even if linked list's size is 0
	l.headSentinel.Next.Prev = newNode
	l.headSentinel.Next = newNode

	l.length++
}

// AddAtTail works in O(1) time, O(1) space
func (l *LinkedList) AddAtTail(val int) {
	newNode := &data_structures.DoublyLinkedListNode[int]{
		Val:  val,
		Prev: l.tailSentinel.Prev,
		Next: l.tailSentinel,
	}

	// We don't need to check for nil, even if linked list's size is 0
	l.tailSentinel.Prev.Next = newNode
	l.tailSentinel.Prev = newNode

	l.length++
}
```

### Slow and fast pointers

They help to find the middle of a linked list and also to check it for cycles (they'll eventually intercept if there are cycles)

```go
func GetMiddleNode(head *data_structures.SinglyLinkedListNode) *data_structures.SinglyLinkedListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
```
```go
func HasCycle(head *data_structures.SinglyLinkedListNode) bool {
	p1 := head
	p2 := head.Next

	for p1 != p2 {
		// We've reached the end
		if p2 == nil || p2.Next == nil {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return true
}
```

### Reverse linked list

Just a small snipped that can be reused. You can reverse a linked list in `O(n)` time, `O(1)` space.
I haven't included it in the [leetcode_library.template](../../build/leetcode_library.template), because I think you should code it quickly yourself.

```go
func ReverseList(head *data_structures.SinglyLinkedListNode[int]) *data_structures.SinglyLinkedListNode[int] {
	var lastReversed *data_structures.SinglyLinkedListNode[int]
	cur := head

	for cur != nil {
		// Save next val
		nextCur := cur.Next

		// Update links
		cur.Next = lastReversed
		lastReversed = cur

		// Iterate to the next
		cur = nextCur
	}

	// New head
	return lastReversed
}
```
