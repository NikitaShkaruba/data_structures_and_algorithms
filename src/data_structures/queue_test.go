package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	assert.Equal(t, 1, q.PeekFirst())
	assert.Equal(t, 1, q.PeekLast())

	q.Enqueue(2)
	assert.Equal(t, 1, q.PeekFirst())
	assert.Equal(t, 2, q.PeekLast())

	q.Dequeue()
	assert.Equal(t, 2, q.PeekFirst())
	assert.Equal(t, 2, q.PeekLast())
}

func TestQueue_GetSize(t *testing.T) {
	q := NewQueue[int]()

	assert.Equal(t, 0, q.GetSize())

	q.Enqueue(1)
	assert.Equal(t, 1, q.GetSize())

	q.Enqueue(1)
	assert.Equal(t, 2, q.GetSize())

	q.Dequeue()
	q.Dequeue()
	assert.Equal(t, 0, q.GetSize())
}

func TestQueue_DifferentGenericTypes(t *testing.T) {
	intStack := NewQueue[int]()
	intStack.Enqueue(1)
	assert.Equal(t, 1, intStack.Dequeue())

	floatStack := NewQueue[float64]()
	floatStack.Enqueue(1.0)
	assert.Equal(t, 1.0, floatStack.Dequeue())

	stringStack := NewQueue[string]()
	stringStack.Enqueue("test")
	assert.Equal(t, "test", stringStack.Dequeue())

	boolStack := NewQueue[bool]()
	boolStack.Enqueue(true)
	assert.Equal(t, true, boolStack.Dequeue())

	type person struct {
		name string
		age  int
	}
	customStructStack := NewQueue[person]()
	customStructStack.Enqueue(person{"test", 123})
	assert.Equal(t, person{"test", 123}, customStructStack.Dequeue())
}
