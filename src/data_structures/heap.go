package data_structures

import (
	"container/heap"
)

////////////////////// Heap //////////////////////

// Heap is a data structure that can retrieve the min/max element in O(logn) time, O(1) space, and insert a new element in O(logn) time, O(1) space
type Heap[T any] struct {
	adapter *stdlibHeapAdapter[T]
}

// NewEmptyHeap creates a new empty heap in O(1) time, O(1) space
// comparator should contain a < b if you want a min heap, and a > b if you want a max heap
func NewEmptyHeap[T any](comparator func(a, b T) bool) *Heap[T] {
	return NewHeapFromArray(make([]T, 0), comparator)
}

// NewHeapFromArray creates a new heap from the given array in O(n) time, O(n) space
// comparator should contain a < b if you want a min heap, and a > b if you want a max heap
func NewHeapFromArray[T any](arr []T, comparator func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		adapter: newHeapAdapterFromArray(arr, comparator),
	}
}

// Push works in O(logn) time, O(1) space
func (h *Heap[T]) Push(val T) {
	heap.Push(h.adapter, val)
}

// Pop works in O(logn) time, O(1) space
func (h *Heap[T]) Pop() T {
	return heap.Pop(h.adapter).(T)
}

// Peek works in O(1) time, O(1) space
func (h *Heap[T]) Peek() T {
	return h.adapter.Peek().(T)
}

// GetSize works in O(1) time, O(1) space
func (h *Heap[T]) GetSize() int {
	return h.adapter.Len()
}

////////////////////// Stdlib Heap Adapter //////////////////////

// stdlibHeapAdapter is needed to not implement the whole heap yourself.
// std library has a semi-convenient (not) interface that does all the dirty work for you.
// I don't like it at all, but it works.
type stdlibHeapAdapter[T any] struct {
	arr        []T
	comparator func(i, j T) bool
}

func newHeapAdapterFromArray[T any](arr []T, comparator func(a, b T) bool) *stdlibHeapAdapter[T] {
	a := &stdlibHeapAdapter[T]{
		arr:        arr,
		comparator: comparator,
	}
	heap.Init(a)

	return a
}

func (a stdlibHeapAdapter[T]) Less(i, j int) bool {
	return a.comparator(a.arr[i], a.arr[j])
}

func (a *stdlibHeapAdapter[T]) Push(x interface{}) {
	a.arr = append(a.arr, x.(T))
}

func (a *stdlibHeapAdapter[T]) Pop() interface{} {
	res := a.arr[len(a.arr)-1]
	a.arr = a.arr[:len(a.arr)-1]
	return res
}

func (a *stdlibHeapAdapter[T]) Peek() interface{} {
	return a.arr[0]
}

func (a stdlibHeapAdapter[T]) Swap(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

func (a stdlibHeapAdapter[T]) Len() int {
	return len(a.arr)
}
