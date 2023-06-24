package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchSimple(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 5}

	assert.Equal(t, -1, BinarySearchSimple(arr, 0))
	assert.Equal(t, 0, BinarySearchSimple(arr, 1))
	assert.Equal(t, 1, BinarySearchSimple(arr, 2))
	assert.Equal(t, 3, BinarySearchSimple(arr, 3))
	assert.Equal(t, 5, BinarySearchSimple(arr, 4))
	assert.Equal(t, 6, BinarySearchSimple(arr, 5))
	assert.Equal(t, -1, BinarySearchSimple(arr, 6))
}

func TestBinarySearchLeftmost(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 5}

	assert.Equal(t, -1, BinarySearchLeftmost(arr, 0))
	assert.Equal(t, 0, BinarySearchLeftmost(arr, 1))
	assert.Equal(t, 1, BinarySearchLeftmost(arr, 2))
	assert.Equal(t, 2, BinarySearchLeftmost(arr, 3))
	assert.Equal(t, 5, BinarySearchLeftmost(arr, 4))
	assert.Equal(t, 6, BinarySearchLeftmost(arr, 5))
	assert.Equal(t, -1, BinarySearchLeftmost(arr, 6))
}

func TestBinarySearchRightmost(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 5}

	assert.Equal(t, -1, BinarySearchRightmost(arr, 0))
	assert.Equal(t, 0, BinarySearchRightmost(arr, 1))
	assert.Equal(t, 1, BinarySearchRightmost(arr, 2))
	assert.Equal(t, 4, BinarySearchRightmost(arr, 3))
	assert.Equal(t, 5, BinarySearchRightmost(arr, 4))
	assert.Equal(t, 6, BinarySearchRightmost(arr, 5))
	assert.Equal(t, -1, BinarySearchRightmost(arr, 6))
}
