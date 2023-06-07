package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDoublyLinkedListNode(t *testing.T) {
	node := NewDoublyLinkedListNode(1)
	assert.Equal(t, 1, node.val)
	assert.Nil(t, node.next)
}

func TestNewSentinelDoublyLinkedListNode(t *testing.T) {
	node := NewSentinelDoublyLinkedListNode[int]()
	assert.Equal(t, 0, node.val)
	assert.Nil(t, node.next)
}

func TestNewDoublyLinkedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	headSentinel, tailSentinel := NewDoublyLinkedList(arr)
	assert.NotNil(t, headSentinel)
	assert.Nil(t, headSentinel.prev)
	assert.NotNil(t, tailSentinel)
	assert.Nil(t, tailSentinel.next)

	cur := headSentinel.next
	for i := 0; i < len(arr); i++ {
		assert.Equal(t, arr[i], cur.val)
		cur = cur.next
	}

	cur = tailSentinel.prev
	for i := len(arr) - 1; i >= 0; i-- {
		assert.Equal(t, arr[i], cur.val)
		cur = cur.prev
	}

	assert.Equal(t, headSentinel, cur)
}

func TestNewDoublyLinkedList_DifferentGenericTypes(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5}
	intHeadSentinel, _ := NewDoublyLinkedList(intArr)
	intCur := intHeadSentinel.next
	for _, num := range intArr {
		assert.Equal(t, num, intCur.val)
		intCur = intCur.next
	}

	floatArr := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	floatHeadSentinel, _ := NewDoublyLinkedList(floatArr)
	floatCur := floatHeadSentinel.next
	for _, num := range floatArr {
		assert.Equal(t, num, floatCur.val)
		floatCur = floatCur.next
	}

	stringArr := []string{"hello", "world", "I'm", "awake"}
	stringHeadSentinel, _ := NewDoublyLinkedList(stringArr)
	stringCur := stringHeadSentinel.next
	for _, str := range stringArr {
		assert.Equal(t, str, stringCur.val)
		stringCur = stringCur.next
	}

	boolArr := []bool{true, false, false, true}
	boolHeadSentinel, _ := NewDoublyLinkedList(boolArr)
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
	customStructHeadSentinel, _ := NewDoublyLinkedList(customStructArr)
	customStructCur := customStructHeadSentinel.next
	for _, str := range customStructArr {
		assert.Equal(t, str, customStructCur.val)
		customStructCur = customStructCur.next
	}
}
