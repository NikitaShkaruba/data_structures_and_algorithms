package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeNode_New(t *testing.T) {
	leftChild := NewLeafBinaryTreeNode(1)
	assert.Equal(t, 1, leftChild.Val)
	assert.Nil(t, leftChild.Left)
	assert.Nil(t, leftChild.Right)

	parent := NewBinaryTreeNode(2, leftChild, nil)
	assert.Equal(t, 2, parent.Val)
	assert.Equal(t, parent.Left, leftChild)
	assert.Nil(t, parent.Right)
}
