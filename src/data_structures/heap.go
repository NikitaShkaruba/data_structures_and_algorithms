package data_structures

import (
	"container/heap"
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	adapter *stdlibHeapAdapter[T]
}

func NewHeap[T constraints.Ordered](isMaxHeap bool) *Heap[T] {
	return NewHeapFromArray(make([]T, 0), isMaxHeap)
}

func NewHeapFromArray[T constraints.Ordered](vals []T, isMaxHeap bool) *Heap[T] {
	return &Heap[T]{
		adapter: newHeapAdapterFromArray(vals, isMaxHeap),
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

////////////////////// stdlib heap adapter //////////////////////

type stdlibHeapAdapter[T constraints.Ordered] struct {
	arr       []T
	isMaxHeap bool
}

func newHeapAdapterFromArray[T constraints.Ordered](vals []T, isMinHeap bool) *stdlibHeapAdapter[T] {
	a := &stdlibHeapAdapter[T]{
		arr:       vals,
		isMaxHeap: isMinHeap,
	}
	heap.Init(a)

	return a
}

func (a stdlibHeapAdapter[T]) Less(i, j int) bool {
	if a.isMaxHeap {
		return a.arr[i] > a.arr[j]
	} else {
		return a.arr[i] < a.arr[j]
	}
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
