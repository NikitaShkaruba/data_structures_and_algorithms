package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedListNode_New(t *testing.T) {
	node := NewDoublyLinkedListNode(
		1,
		&DoublyLinkedListNode[int]{
			Val: 2,
		},
		&DoublyLinkedListNode[int]{
			Val: 3,
		},
	)

	assert.Equal(t, 1, node.Val)
	assert.Equal(t, 2, node.Prev.Val)
	assert.Equal(t, 3, node.Next.Val)
}

func TestDoublyLinkedListNode_NewSentinel(t *testing.T) {
	node := NewSentinelDoublyLinkedListNode[int]()

	assert.Equal(t, 0, node.Val)
	assert.Nil(t, node.Prev)
	assert.Nil(t, node.Next)
}

func TestDoublyLinkedList_New(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	headSentinel, tailSentinel := NewDoublyLinkedList(arr)
	assert.NotNil(t, headSentinel)
	assert.Nil(t, headSentinel.Prev)
	assert.NotNil(t, tailSentinel)
	assert.Nil(t, tailSentinel.Next)

	cur := headSentinel.Next
	for i := 0; i < len(arr); i++ {
		assert.Equal(t, arr[i], cur.Val)
		cur = cur.Next
	}

	cur = tailSentinel.Prev
	for i := len(arr) - 1; i >= 0; i-- {
		assert.Equal(t, arr[i], cur.Val)
		cur = cur.Prev
	}

	assert.Equal(t, headSentinel, cur)
}

func TestDoublyLinkedList_DifferentGenericTypes(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5}
	intHeadSentinel, _ := NewDoublyLinkedList(intArr)
	intCur := intHeadSentinel.Next
	for _, num := range intArr {
		assert.Equal(t, num, intCur.Val)
		intCur = intCur.Next
	}

	floatArr := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	floatHeadSentinel, _ := NewDoublyLinkedList(floatArr)
	floatCur := floatHeadSentinel.Next
	for _, num := range floatArr {
		assert.Equal(t, num, floatCur.Val)
		floatCur = floatCur.Next
	}

	stringArr := []string{"hello", "world", "I'm", "awake"}
	stringHeadSentinel, _ := NewDoublyLinkedList(stringArr)
	stringCur := stringHeadSentinel.Next
	for _, str := range stringArr {
		assert.Equal(t, str, stringCur.Val)
		stringCur = stringCur.Next
	}

	boolArr := []bool{true, false, false, true}
	boolHeadSentinel, _ := NewDoublyLinkedList(boolArr)
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
	customStructHeadSentinel, _ := NewDoublyLinkedList(customStructArr)
	customStructCur := customStructHeadSentinel.Next
	for _, str := range customStructArr {
		assert.Equal(t, str, customStructCur.Val)
		customStructCur = customStructCur.Next
	}
}
