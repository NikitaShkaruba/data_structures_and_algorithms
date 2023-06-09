package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_NewEmpty(t *testing.T) {
	q := NewEmptyQueue[int]()
	assert.Equal(t, 0, q.GetSize())
}

func TestQueue_NewFromArray(t *testing.T) {
	d := NewQueueFromArray([]int{1, 2, 3})

	assert.Equal(t, 3, d.GetSize())

	assert.Equal(t, 1, d.Dequeue())
	assert.Equal(t, 2, d.Dequeue())
	assert.Equal(t, 3, d.Dequeue())
}

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := NewEmptyQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
}

func TestQueue_Peek(t *testing.T) {
	q := NewEmptyQueue[int]()

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
	q := NewEmptyQueue[int]()

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
	intStack := NewEmptyQueue[int]()
	intStack.Enqueue(1)
	assert.Equal(t, 1, intStack.Dequeue())

	floatStack := NewEmptyQueue[float64]()
	floatStack.Enqueue(1.0)
	assert.Equal(t, 1.0, floatStack.Dequeue())

	stringStack := NewEmptyQueue[string]()
	stringStack.Enqueue("test")
	assert.Equal(t, "test", stringStack.Dequeue())

	boolStack := NewEmptyQueue[bool]()
	boolStack.Enqueue(true)
	assert.Equal(t, true, boolStack.Dequeue())

	type person struct {
		name string
		age  int
	}
	customStructStack := NewEmptyQueue[person]()
	customStructStack.Enqueue(person{"test", 123})
	assert.Equal(t, person{"test", 123}, customStructStack.Dequeue())
}
