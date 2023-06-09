package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexedHeap_NewEmpty(t *testing.T) {
	h := NewEmptyIndexedHeap[int](func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 0, h.GetSize())
}

func TestIndexedHeap_NewFromArray(t *testing.T) {
	arr := []int{12, 4, 8}
	h := NewIndexedHeapFromArray[int](arr, func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 12, h.Pop())
}

func TestIndexedHeap_PushPopMax(t *testing.T) {
	h := NewEmptyIndexedHeap[int](func(a, b int) bool {
		return a > b
	})

	h.Push(12)
	h.Push(4)
	h.Push(8)

	assert.Equal(t, 12, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 4, h.Pop())
}

func TestIndexedHeap_PushPopMin(t *testing.T) {
	h := NewEmptyIndexedHeap[int](func(a, b int) bool {
		return a < b
	})

	h.Push(12)
	h.Push(4)
	h.Push(8)

	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 12, h.Pop())
}

func TestIndexedHeap_Peek(t *testing.T) {
	h := NewIndexedHeapFromArray([]int{12, 4, 8}, func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 4, h.Peek())
	h.Pop()

	assert.Equal(t, 8, h.Peek())
	h.Pop()

	assert.Equal(t, 12, h.Peek())
	h.Pop()
}

func TestIndexedHeap_RemoveByValue(t *testing.T) {
	h := NewIndexedHeapFromArray([]int{12, 4, 4, 8}, func(a, b int) bool {
		return a < b
	})

	assert.False(t, h.RemoveByValue(100000))
	assert.Equal(t, 4, h.GetSize())

	assert.True(t, h.RemoveByValue(4))
	assert.Equal(t, 4, h.Peek())
	assert.Equal(t, 3, h.GetSize())

	assert.True(t, h.RemoveByValue(4))
	assert.Equal(t, 8, h.Peek())
	assert.Equal(t, 2, h.GetSize())

	assert.True(t, h.RemoveByValue(12))
	assert.Equal(t, 8, h.Peek())
	assert.Equal(t, 1, h.GetSize())

	assert.True(t, h.RemoveByValue(8))
	assert.Equal(t, 0, h.GetSize())
}

func TestIndexedHeap_GetSize(t *testing.T) {
	h := NewEmptyIndexedHeap[int](func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 0, h.GetSize())

	h.Push(12)
	assert.Equal(t, 1, h.GetSize())

	h.Push(4)
	assert.Equal(t, 2, h.GetSize())

	h.Pop()
	assert.Equal(t, 1, h.GetSize())

	h.Pop()
	assert.Equal(t, 0, h.GetSize())
}

func TestIndexedHeap_DifferentGenericTypes(t *testing.T) {
	intHeap := NewIndexedHeapFromArray([]int{12, 4, 8}, func(a, b int) bool {
		return a < b
	})
	assert.Equal(t, 4, intHeap.Pop())
	assert.Equal(t, 8, intHeap.Pop())
	assert.Equal(t, 12, intHeap.Pop())

	floatHeap := NewIndexedHeapFromArray([]float64{12.0, 4.0, 8.0}, func(a, b float64) bool {
		return a < b
	})
	assert.Equal(t, 4.0, floatHeap.Pop())
	assert.Equal(t, 8.0, floatHeap.Pop())
	assert.Equal(t, 12.0, floatHeap.Pop())

	stringHeap := NewIndexedHeapFromArray([]string{"aaa", "a", "aa"}, func(a, b string) bool {
		return a < b
	})
	assert.Equal(t, "a", stringHeap.Pop())
	assert.Equal(t, "aa", stringHeap.Pop())
	assert.Equal(t, "aaa", stringHeap.Pop())

	type person struct {
		name string
		age  int
	}
	people := []person{
		{name: "dave", age: 12},
		{name: "johny", age: 4},
		{name: "danny", age: 8},
	}
	customStructHeap := NewIndexedHeapFromArray(people, func(a, b person) bool {
		return a.age < b.age
	})
	assert.Equal(t, "johny", customStructHeap.Pop().name)
	assert.Equal(t, "danny", customStructHeap.Pop().name)
	assert.Equal(t, "dave", customStructHeap.Pop().name)
}
