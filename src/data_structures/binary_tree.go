package data_structures

////////////////////// BinaryTreeNode //////////////////////
// You can find this data structure overview in docs/data_structures/binary_tree.md

// BinaryTreeNode is a node with value and pointers to the left and the right child
type BinaryTreeNode[T any] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

// NewBinaryTreeNode works in O(1) time, O(1) space
func NewBinaryTreeNode[T any](val T, left, right *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

// NewLeafBinaryTreeNode works in O(1) time, O(1) space
func NewLeafBinaryTreeNode[T any](val T) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{
		Val: val,
	}
}
