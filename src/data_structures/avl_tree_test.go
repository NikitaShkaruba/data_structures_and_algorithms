package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvlTree_New(t *testing.T) {
	tree := NewEmptyAvlTree[int](minIntTreeComparator)

	assert.Equal(t, 0, tree.GetSize())
}

func TestAvlTree_NewFromArray(t *testing.T) {
	arr := []int{8, 19, 4}
	tree := NewAvlTreeFromArray[int](arr, minIntTreeComparator)

	assert.Equal(t, 3, tree.GetSize())
	assert.Equal(t, 19, tree.GetMax())
	assert.Equal(t, 4, tree.GetMin())
}

func TestAvlTree_InsertDeleteGetMinGetMax(t *testing.T) {
	tree := NewEmptyAvlTree[int](minIntTreeComparator)

	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(4)
	assert.Equal(t, 4, tree.GetMin())
	assert.Equal(t, 12, tree.GetMax())

	tree.Insert(2)
	tree.Insert(22)
	assert.Equal(t, 2, tree.GetMin())
	assert.Equal(t, 22, tree.GetMax())

	tree.Insert(10)
	tree.Insert(21)
	assert.Equal(t, 2, tree.GetMin())
	assert.Equal(t, 22, tree.GetMax())

	assert.False(t, tree.Contains(123123))
	assert.True(t, tree.Contains(8))

	tree.Delete(22)
	assert.Equal(t, 2, tree.GetMin())
	assert.Equal(t, 21, tree.GetMax())
	assert.False(t, tree.Contains(22))
}

func TestAvlTree_InsertDeleteGetMinGetMaxOnMaxTree(t *testing.T) {
	tree := NewEmptyAvlTree[int](func(a int, b int) int {
		return -minIntTreeComparator(a, b)
	})

	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(4)
	assert.Equal(t, 12, tree.GetMin())
	assert.Equal(t, 4, tree.GetMax())

	tree.Insert(2)
	tree.Insert(22)
	assert.Equal(t, 22, tree.GetMin())
	assert.Equal(t, 2, tree.GetMax())

	tree.Insert(10)
	tree.Insert(21)
	assert.Equal(t, 22, tree.GetMin())
	assert.Equal(t, 2, tree.GetMax())

	assert.False(t, tree.Contains(123123))
	assert.True(t, tree.Contains(8))

	tree.Delete(22)
	assert.Equal(t, 21, tree.GetMin())
	assert.Equal(t, 2, tree.GetMax())
	assert.False(t, tree.Contains(22))
}

func TestAvlTree_GetSize(t *testing.T) {
	tree := NewEmptyAvlTree[int](minIntTreeComparator)

	tree.Insert(1)
	assert.Equal(t, 1, tree.GetSize())

	tree.Insert(2)
	assert.Equal(t, 2, tree.GetSize())

	tree.Delete(2)
	assert.Equal(t, 1, tree.GetSize())

	tree.Delete(1)
	assert.Equal(t, 0, tree.GetSize())
}

func TestAvlTree_DifferentGenericTypes(t *testing.T) {
	intTree := NewAvlTreeFromArray([]int{12, 4, 8}, minIntTreeComparator)
	assert.Equal(t, 4, intTree.GetMin())
	assert.Equal(t, 12, intTree.GetMax())

	floatTree := NewAvlTreeFromArray([]float64{12.0, 4.0, 8.0}, func(a, b float64) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	})
	assert.Equal(t, 4.0, floatTree.GetMin())
	assert.Equal(t, 12.0, floatTree.GetMax())

	stringTree := NewAvlTreeFromArray([]string{"aaa", "a", "aa"}, func(a, b string) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	})
	assert.Equal(t, "a", stringTree.GetMin())
	assert.Equal(t, "aaa", stringTree.GetMax())

	type person struct {
		name string
		age  int
	}
	people := []person{
		{name: "dave", age: 12},
		{name: "johny", age: 4},
		{name: "danny", age: 8},
	}
	customStructTree := NewAvlTreeFromArray(people, func(a, b person) int {
		if a.age > b.age {
			return 1
		} else if a.age < b.age {
			return -1
		} else {
			return 0
		}
	})
	assert.Equal(t, "johny", customStructTree.GetMin().name)
	assert.Equal(t, "dave", customStructTree.GetMax().name)
}

////////////////////// Helpers //////////////////////

func minIntTreeComparator(a int, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
