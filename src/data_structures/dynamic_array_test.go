package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicArray_PushBack(t *testing.T) {
	a := NewEmptyDynamicArray[int]()
	assert.Equal(t, 0, a.GetSize())
	assert.Equal(t, 2, a.GetCapacity())

	a.PushBack(0)
	assert.Equal(t, 1, a.GetSize())
	assert.Equal(t, 2, a.GetCapacity())

	a.PushBack(1)
	assert.Equal(t, 2, a.GetSize())
	assert.Equal(t, 2, a.GetCapacity())

	a.PushBack(2)
	assert.Equal(t, 3, a.GetSize())
	assert.Equal(t, 4, a.GetCapacity())

	a.PushBack(3)
	assert.Equal(t, 4, a.GetSize())
	assert.Equal(t, 4, a.GetCapacity())

	a.PushBack(4)
	assert.Equal(t, 5, a.GetSize())
	assert.Equal(t, 8, a.GetCapacity())
}

func TestDynamicArray_PopBack(t *testing.T) {
	a := NewEmptyDynamicArray[int]()
	a.PushBack(0)
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)
	assert.Equal(t, 5, a.GetSize())
	assert.Equal(t, 8, a.GetCapacity())

	lastVal := a.PopBack()
	assert.Equal(t, 4, lastVal)
	assert.Equal(t, 4, a.GetSize())
	assert.Equal(t, 8, a.GetCapacity())

	lastVal = a.PopBack()
	assert.Equal(t, 3, lastVal)
	assert.Equal(t, 3, a.GetSize())
	assert.Equal(t, 8, a.GetCapacity())

	lastVal = a.PopBack()
	assert.Equal(t, 2, lastVal)
	assert.Equal(t, 2, a.GetSize())
	assert.Equal(t, 8, a.GetCapacity())

	lastVal = a.PopBack()
	assert.Equal(t, 1, lastVal)
	assert.Equal(t, 1, a.GetSize())
	assert.Equal(t, 4, a.GetCapacity())

	lastVal = a.PopBack()
	assert.Equal(t, 0, lastVal)
	assert.Equal(t, 0, a.GetSize())
	assert.Equal(t, 2, a.GetCapacity())
}

func TestDynamicArray_GetSet(t *testing.T) {
	a := NewEmptyDynamicArray[int]()
	a.PushBack(0)
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)

	assert.Equal(t, 3, a.Get(3))
	a.Set(3, 33)
	assert.Equal(t, 33, a.Get(3))

	assert.Equal(t, 0, a.Get(0))
	a.Set(0, -12)
	assert.Equal(t, -12, a.Get(0))
}
