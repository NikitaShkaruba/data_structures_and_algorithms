package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglyLinkedListNode_New(t *testing.T) {
	firstNode := NewSinglyLinkedListNode(1, nil)
	assert.Equal(t, 1, firstNode.Val)
	assert.Nil(t, firstNode.Next)

	secondNode := NewSinglyLinkedListNode(2, firstNode)
	assert.Equal(t, 2, secondNode.Val)
	assert.Equal(t, firstNode, secondNode.Next)
}

func TestSinglyLinkedListNode_NewSentinel(t *testing.T) {
	node := NewSentinelSinglyLinkedListNode[int]()
	assert.Equal(t, 0, node.Val)
	assert.Nil(t, node.Next)
}

func TestSinglyLinkedList_New(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	headSentinel := NewSinglyLinkedList(arr)
	assert.NotNil(t, headSentinel)

	cur := headSentinel.Next
	for _, num := range arr {
		assert.Equal(t, num, cur.Val)
		cur = cur.Next
	}
}

func TestSinglyLinkedList_DifferentGenericTypes(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5}
	intHeadSentinel := NewSinglyLinkedList(intArr)
	intCur := intHeadSentinel.Next
	for _, num := range intArr {
		assert.Equal(t, num, intCur.Val)
		intCur = intCur.Next
	}

	floatArr := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	floatHeadSentinel := NewSinglyLinkedList(floatArr)
	floatCur := floatHeadSentinel.Next
	for _, num := range floatArr {
		assert.Equal(t, num, floatCur.Val)
		floatCur = floatCur.Next
	}

	stringArr := []string{"hello", "world", "I'm", "awake"}
	stringHeadSentinel := NewSinglyLinkedList(stringArr)
	stringCur := stringHeadSentinel.Next
	for _, str := range stringArr {
		assert.Equal(t, str, stringCur.Val)
		stringCur = stringCur.Next
	}

	boolArr := []bool{true, false, false, true}
	boolHeadSentinel := NewSinglyLinkedList(boolArr)
	boolCur := boolHeadSentinel.Next
	for _, str := range boolArr {
		assert.Equal(t, str, boolCur.Val)
		boolCur = boolCur.Next
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
	customStructCur := customStructHeadSentinel.Next
	for _, str := range customStructArr {
		assert.Equal(t, str, customStructCur.Val)
		customStructCur = customStructCur.Next
	}
}
