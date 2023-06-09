package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeque_NewEmpty(t *testing.T) {
	d := NewEmptyDeque[int]()
	assert.Equal(t, 0, d.GetSize())
}

func TestDeque_NewFromArray(t *testing.T) {
	d := NewDequeFromArray([]int{1, 2, 3})

	assert.Equal(t, 3, d.GetSize())

	assert.Equal(t, 1, d.PopFront())
	assert.Equal(t, 2, d.PopFront())
	assert.Equal(t, 3, d.PopFront())
}

func TestDeque_PushPop(t *testing.T) {
	d := NewEmptyDeque[int]()

	d.PushFront(2)
	d.PushFront(1)
	d.PushBack(3)
	d.PushBack(4)

	assert.Equal(t, 4, d.PopBack())
	assert.Equal(t, 3, d.PopBack())
	assert.Equal(t, 1, d.PopFront())
	assert.Equal(t, 2, d.PopFront())
}

func TestDeque_Peek(t *testing.T) {
	d := NewEmptyDeque[int]()

	d.PushFront(2)
	assert.Equal(t, 2, d.PeekFront())
	assert.Equal(t, 2, d.PeekBack())

	d.PushFront(1)
	assert.Equal(t, 1, d.PeekFront())
	assert.Equal(t, 2, d.PeekBack())

	d.PushBack(3)
	assert.Equal(t, 1, d.PeekFront())
	assert.Equal(t, 3, d.PeekBack())

	d.PushBack(4)
	assert.Equal(t, 1, d.PeekFront())
	assert.Equal(t, 4, d.PeekBack())
}

func TestDeque_GetSize(t *testing.T) {
	d := NewEmptyDeque[int]()

	d.PushFront(2)
	assert.Equal(t, 1, d.GetSize())

	d.PushBack(1)
	assert.Equal(t, 2, d.GetSize())

	d.PopFront()
	assert.Equal(t, 1, d.GetSize())

	d.PopBack()
	assert.Equal(t, 0, d.GetSize())
}

func TestDeque_DifferentGenericTypes(t *testing.T) {
	intDeque := NewEmptyDeque[int]()
	intDeque.PushFront(1)
	assert.Equal(t, 1, intDeque.PopFront())

	floatDeque := NewEmptyDeque[float64]()
	floatDeque.PushFront(1.0)
	assert.Equal(t, 1.0, floatDeque.PopFront())

	stringDeque := NewEmptyDeque[string]()
	stringDeque.PushFront("test")
	assert.Equal(t, "test", stringDeque.PopFront())

	boolDeque := NewEmptyDeque[bool]()
	boolDeque.PushFront(true)
	assert.Equal(t, true, boolDeque.PopFront())

	type person struct {
		name string
		age  int
	}
	customStructDeque := NewEmptyDeque[person]()
	customStructDeque.PushFront(person{"test", 123})
	assert.Equal(t, person{"test", 123}, customStructDeque.PopFront())
}
