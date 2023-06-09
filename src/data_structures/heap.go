package data_structures

import (
	"container/heap"
)

////////////////////// Heap //////////////////////

type Heap[T any] struct {
	adapter *stdlibHeapAdapter[T]
}

// NewHeap creates a new empty heap.
// comparator should contain a < b if you want a min heap, and a > b if you want a max heap
func NewHeap[T any](comparator func(a, b T) bool) *Heap[T] {
	return NewHeapFromArray(make([]T, 0), comparator)
}

// NewHeapFromArray creates a new heap from the given array.
// comparator should contain a < b if you want a min heap, and a > b if you want a max heap
func NewHeapFromArray[T any](vals []T, comparator func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		adapter: newHeapAdapterFromArray(vals, comparator),
	}
}

func (h *Heap[T]) Push(val T) {
	heap.Push(h.adapter, val)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.adapter).(T)
}

func (h *Heap[T]) GetSize() int {
	return h.adapter.Len()
}

////////////////////// Stdlib Heap Adapter //////////////////////

type stdlibHeapAdapter[T any] struct {
	arr        []T
	comparator func(i, j T) bool
}

func newHeapAdapterFromArray[T any](vals []T, comparator func(a, b T) bool) *stdlibHeapAdapter[T] {
	a := &stdlibHeapAdapter[T]{
		arr:        vals,
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
	res := (a.arr)[len(a.arr)-1]
	a.arr = (a.arr)[:len(a.arr)-1]
	return res
}

func (a *stdlibHeapAdapter[T]) Peek() interface{} {
	return (a.arr)[0]
}

func (a stdlibHeapAdapter[T]) Swap(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

func (a stdlibHeapAdapter[T]) Len() int {
	return len(a.arr)
}
