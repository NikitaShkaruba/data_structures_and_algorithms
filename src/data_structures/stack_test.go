package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	s := NewEmptyStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())
}

func TestStack_Peek(t *testing.T) {
	s := NewEmptyStack[int]()

	s.Push(1)
	assert.Equal(t, 1, s.Peek())

	s.Push(2)
	assert.Equal(t, 2, s.Peek())

	s.Push(3)
	assert.Equal(t, 3, s.Peek())
}

func TestStack_GetSize(t *testing.T) {
	s := NewEmptyStack[int]()

	s.Push(1)
	assert.Equal(t, 1, s.GetSize())

	s.Push(1)
	assert.Equal(t, 2, s.GetSize())

	s.Pop()
	assert.Equal(t, 1, s.GetSize())

	s.Pop()
	assert.Equal(t, 0, s.GetSize())
}

func TestStack_DifferentGenericTypes(t *testing.T) {
	intStack := NewEmptyStack[int]()
	intStack.Push(1)
	assert.Equal(t, 1, intStack.Pop())

	floatStack := NewEmptyStack[float64]()
	floatStack.Push(1.0)
	assert.Equal(t, 1.0, floatStack.Pop())

	stringStack := NewEmptyStack[string]()
	stringStack.Push("test")
	assert.Equal(t, "test", stringStack.Pop())

	boolStack := NewEmptyStack[bool]()
	boolStack.Push(true)
	assert.Equal(t, true, boolStack.Pop())

	type person struct {
		name string
		age  int
	}
	customStructStack := NewEmptyStack[person]()
	customStructStack.Push(person{"test", 123})
	assert.Equal(t, person{"test", 123}, customStructStack.Pop())
}
