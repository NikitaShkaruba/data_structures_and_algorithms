package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSinglyLinkedListNode(t *testing.T) {
	node := NewSinglyLinkedListNode(1)
	assert.Equal(t, 1, node.val)
	assert.Nil(t, node.next)
}

func TestNewSentinelSinglyLinkedListNode(t *testing.T) {
	node := NewSentinelSinglyLinkedListNode[int]()
	assert.Equal(t, 0, node.val)
	assert.Nil(t, node.next)
}

func TestNewSinglyLinkedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	headSentinel := NewSinglyLinkedList(arr)
	assert.NotNil(t, headSentinel)

	cur := headSentinel.next
	for _, num := range arr {
		assert.Equal(t, num, cur.val)
		cur = cur.next
	}
}

func TestNewSinglyLinkedList_DifferentGenericTypes(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5}
	intHeadSentinel := NewSinglyLinkedList(intArr)
	intCur := intHeadSentinel.next
	for _, num := range intArr {
		assert.Equal(t, num, intCur.val)
		intCur = intCur.next
	}

	floatArr := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	floatHeadSentinel := NewSinglyLinkedList(floatArr)
	floatCur := floatHeadSentinel.next
	for _, num := range floatArr {
		assert.Equal(t, num, floatCur.val)
		floatCur = floatCur.next
	}

	stringArr := []string{"hello", "world", "I'm", "awake"}
	stringHeadSentinel := NewSinglyLinkedList(stringArr)
	stringCur := stringHeadSentinel.next
	for _, str := range stringArr {
		assert.Equal(t, str, stringCur.val)
		stringCur = stringCur.next
	}

	boolArr := []bool{true, false, false, true}
	boolHeadSentinel := NewSinglyLinkedList(boolArr)
	boolCur := boolHeadSentinel.next
	for _, str := range boolArr {
		assert.Equal(t, str, boolCur.val)
		boolCur = boolCur.next
	}

	type person struct {
		name string
		age  int
	}
	customStructArr := []person{
		{name: "dave", age: 23},
		{name: "johny", age: 30},
	}
	customStructHeadSentinel := NewSinglyLinkedList(customStructArr)
	customStructCur := customStructHeadSentinel.next
	for _, str := range customStructArr {
		assert.Equal(t, str, customStructCur.val)
		customStructCur = customStructCur.next
	}
}
