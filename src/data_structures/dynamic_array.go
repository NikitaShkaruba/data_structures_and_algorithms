package data_structures

////////////////////// Dynamic array //////////////////////
// @DONT_INCLUDE_IN_TEMPLATE
// You can find this data structure overview in docs/data_structures/dynamic_array.md

type DynamicArray[T any] struct {
	staticArray []T
	size        int
	capacity    int
}

func NewEmptyDynamicArray[T any]() DynamicArray[T] {
	return DynamicArray[T]{
		staticArray: make([]T, 2),
		size:        0,
		capacity:    2,
	}
}

// PushBack works in O(1) time, O(1) space
func (a *DynamicArray[T]) PushBack(val T) {
	if a.size == a.capacity {
		a.expand()
	}

	a.staticArray[a.size] = val
	a.size++
}

// PopBack works in O(1) time, O(1) space
func (a *DynamicArray[T]) PopBack() T {
	if a.size == 0 {
		panic("no elements left")
	}

	a.size--
	if a.size < a.capacity/3 {
		a.shrink()
	}

	return a.staticArray[a.size]
}

// Get works in O(1) time, O(1) space
func (a *DynamicArray[T]) Get(i int) T {
	if i < 0 || i >= a.size {
		panic("element doesn't exist")
	}

	return a.staticArray[i]
}

// Set works in O(1) time, O(1) space
func (a *DynamicArray[T]) Set(i int, val T) {
	if i < 0 || i >= a.size {
		panic("element doesn't exist")
	}

	a.staticArray[i] = val
}

// GetCapacity works in O(1) time, O(1) space
func (a *DynamicArray[T]) GetCapacity() int {
	return a.capacity
}

// GetSize works in O(1) time, O(1) space
func (a *DynamicArray[T]) GetSize() int {
	return a.size
}

func (a *DynamicArray[T]) expand() {
	a.capacity *= 2

	oldStaticArray := a.staticArray

	a.staticArray = make([]T, a.capacity)
	for i := 0; i < a.size; i++ {
		a.staticArray[i] = oldStaticArray[i]
	}
}

func (a *DynamicArray[T]) shrink() {
	a.capacity /= 2

	oldStaticArray := a.staticArray

	a.staticArray = make([]T, a.capacity)
	for i := 0; i <= a.size; i++ {
		a.staticArray[i] = oldStaticArray[i]
	}
}
