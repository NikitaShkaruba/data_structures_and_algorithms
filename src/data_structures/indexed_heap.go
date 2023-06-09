package data_structures

import "container/heap"

//////////////////////// Indexed Heap ////////////////////////

// IndexedHeap is almost identical to Heap, but allows to remove elements by value for O(logn) time, O(1) space.
// Also, it doesn't spend memory for storing duplicate values.
type IndexedHeap[T comparable] struct {
	adapter     *stdlibIndexedHeapAdapter[T]
	valueCounts map[T]int
	size        int
}

// NewEmptyIndexedHeap works in O(1) time, O(1) space
func NewEmptyIndexedHeap[T comparable](comparator func(a, b T) bool) *IndexedHeap[T] {
	return NewIndexedHeapFromArray(make([]T, 0), comparator)
}

// NewIndexedHeapFromArray works in O(n) time, O(n) space
func NewIndexedHeapFromArray[T comparable](arr []T, comparator func(a, b T) bool) *IndexedHeap[T] {
	uniqueValues := make([]T, 0)
	valueCounts := make(map[T]int)
	for _, n := range arr {
		if valueCounts[n] == 0 {
			uniqueValues = append(uniqueValues, n)
		}
		valueCounts[n]++
	}

	return &IndexedHeap[T]{
		adapter:     newStdlibIndexedHeapAdapter(uniqueValues, comparator),
		valueCounts: valueCounts,
		size:        len(arr),
	}
}

// Push works in O(logn) time, O(1) space
func (h *IndexedHeap[T]) Push(val T) {
	if h.valueCounts[val] == 0 {
		heap.Push(h.adapter, val)
	}

	h.valueCounts[val]++
	h.size++
}

// Pop works in O(logn) time, O(1) space
func (h *IndexedHeap[T]) Pop() T {
	val := h.adapter.Peek().(T)

	h.valueCounts[val]--
	if h.valueCounts[val] == 0 {
		heap.Pop(h.adapter)
		delete(h.valueCounts, val)
	}

	h.size--

	return val
}

// Peek works in O(logn) time, O(1) space
func (h *IndexedHeap[T]) Peek() T {
	return h.adapter.Peek().(T)
}

// RemoveByValue works in O(logn) time, O(1) space
func (h *IndexedHeap[T]) RemoveByValue(val T) bool {
	if h.valueCounts[val] == 0 {
		return false
	}

	h.valueCounts[val]--
	if h.valueCounts[val] == 0 {
		h.adapter.removeByValue(val)
		delete(h.valueCounts, val)
	}

	h.size--

	return true
}

// GetSize works in O(1) time, O(1) space
func (h *IndexedHeap[T]) GetSize() int {
	return h.size
}

//////////////////////// Stdlib Indexed Heap Adapter ////////////////////////

// stdlibIndexedHeapAdapter is needed to not implement the whole heap yourself.
// std library has a semi-convenient (not) interface that does all the dirty work for you.
// I don't like it at all, but it works.
type stdlibIndexedHeapAdapter[T comparable] struct {
	values       []T
	valueIndexes map[T]int
	comparator   func(i, j T) bool
}

func newStdlibIndexedHeapAdapter[T comparable](arr []T, comparator func(a, b T) bool) *stdlibIndexedHeapAdapter[T] {
	valueIndexes := make(map[T]int)
	for i := range arr {
		valueIndexes[arr[i]] = i
	}

	a := &stdlibIndexedHeapAdapter[T]{
		values:       arr,
		valueIndexes: valueIndexes,
		comparator:   comparator,
	}
	heap.Init(a)

	return a
}

func (a stdlibIndexedHeapAdapter[T]) Less(i, j int) bool {
	return a.comparator(a.values[i], a.values[j])
}

func (a *stdlibIndexedHeapAdapter[T]) Push(x interface{}) {
	val := x.(T)

	a.values = append(a.values, val)
	a.valueIndexes[val] = len(a.values) - 1
}

func (a *stdlibIndexedHeapAdapter[T]) Pop() interface{} {
	val := (a.values)[len(a.values)-1]

	a.values = a.values[:len(a.values)-1]
	delete(a.valueIndexes, val)

	return val
}

func (a *stdlibIndexedHeapAdapter[T]) Peek() interface{} {
	return a.values[0]
}

func (a stdlibIndexedHeapAdapter[T]) Swap(i, j int) {
	a.valueIndexes[a.values[i]], a.valueIndexes[a.values[j]] = a.valueIndexes[a.values[j]], a.valueIndexes[a.values[i]]
	a.values[i], a.values[j] = a.values[j], a.values[i]
}

func (a stdlibIndexedHeapAdapter[T]) Len() int {
	return len(a.values)
}

func (a *stdlibIndexedHeapAdapter[T]) removeByValue(val T) {
	i, ok := a.valueIndexes[val]
	if !ok {
		return
	}

	heap.Remove(a, i)
}
