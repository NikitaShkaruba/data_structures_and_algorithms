package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap_PushPopMax(t *testing.T) {
	h := NewHeap[int](true)

	h.Push(12)
	h.Push(4)
	h.Push(8)

	assert.Equal(t, 12, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 4, h.Pop())
}

func TestHeap_PushPopMin(t *testing.T) {
	h := NewHeap[int](false)

	h.Push(12)
	h.Push(4)
	h.Push(8)

	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 12, h.Pop())
}

func TestHeap_GetSize(t *testing.T) {
	h := NewHeap[int](false)

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

func TestHeap_NewFromArray(t *testing.T) {
	h := NewHeapFromArray[int]([]int{12, 4, 8}, false)

	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 12, h.Pop())
}

func TestHeap_DifferentGenericTypes(t *testing.T) {
	intHeap := NewHeapFromArray([]int{12, 4, 8}, false)
	assert.Equal(t, 4, intHeap.Pop())
	assert.Equal(t, 8, intHeap.Pop())
	assert.Equal(t, 12, intHeap.Pop())

	floatHeap := NewHeapFromArray([]float64{12.0, 4.0, 8.0}, false)
	assert.Equal(t, 4.0, floatHeap.Pop())
	assert.Equal(t, 8.0, floatHeap.Pop())
	assert.Equal(t, 12.0, floatHeap.Pop())

	stringHeap := NewHeapFromArray([]string{"aaa", "a", "aa"}, false)
	assert.Equal(t, "a", stringHeap.Pop())
	assert.Equal(t, "aa", stringHeap.Pop())
	assert.Equal(t, "aaa", stringHeap.Pop())

	// TODO: support
	//type person struct {
	//	name string
	//	age  int
	//}
	//customStructHeap := NewHeapFromArray([]person{
	//	{name: "dave", age: 12},
	//	{name: "johny", age: 4},
	//	{name: "danny", age: 8},
	//}, false)
	//assert.Equal(t, "johny", customStructHeap.Pop().name)
	//assert.Equal(t, "danny", customStructHeap.Pop().name)
	//assert.Equal(t, "dave", customStructHeap.Pop().name)
}
