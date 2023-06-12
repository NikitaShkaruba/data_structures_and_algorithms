package data_structures

// This code was copied from the internet and refactored, because of that it's not of the highest quality

////////////////////// AVL tree //////////////////////

// AvlTree is a data structure that can retrieve the min/max element in O(logn) time, O(1) space, and insert a new element in O(logn) time, O(1) space
type AvlTree[T comparable] struct {
	root       *AvlTreeNode[T]
	values     []T
	size       int
	comparator func(a, b T) int
}

// NewEmptyAvlTree works in O(1) time, O(1) space
// lessComparator should return 1 if a > b, -1 if a < b, and 0 if a == b
func NewEmptyAvlTree[T comparable](lessComparator func(a, b T) int) *AvlTree[T] {
	return NewAvlTreeFromArray(make([]T, 0), lessComparator)
}

// NewAvlTreeFromArray works in O(n*logn) time, O(n) space
// lessComparator should return 1 if a > b, -1 if a < b, and 0 if a == b
func NewAvlTreeFromArray[T comparable](arr []T, lessComparator func(a, b T) int) *AvlTree[T] {
	t := &AvlTree[T]{
		comparator: lessComparator,
	}

	for _, val := range arr {
		t.Insert(val)
	}

	return t
}

// Insert works in O(logn) time, O(n) space
func (t *AvlTree[T]) Insert(value T) {
	added := false
	t.root = insertNode(t.root, value, &added, t.comparator)
	if added {
		t.size++
	}
	t.values = nil
}

// Delete works in O(logn) time, O(n) space
func (t *AvlTree[T]) Delete(value T) {
	deleted := false
	t.root = deleteNode(t.root, value, &deleted)
	if deleted {
		t.size--
	}
	t.values = nil
}

// Contains works in O(logn) time, O(n) space
func (t *AvlTree[T]) Contains(value T) bool {
	var node *AvlTreeNode[T]

	if t.root != nil {
		node = t.root.get(value)
	}

	if node == nil {
		return false
	}

	return true
}

// GetMin works in O(logn) time, O(n) space
func (t *AvlTree[T]) GetMin() T {
	if t.root == nil {
		return *new(T)
	}

	var beginning = t.root
	for beginning.left != nil {
		beginning = beginning.left
	}

	if beginning == nil {
		for beginning.right != nil {
			beginning = beginning.right
		}
	}

	if beginning == nil {
		return *new(T)
	}

	return beginning.Value
}

// GetMax works in O(logn) time, O(n) space
func (t *AvlTree[T]) GetMax() T {
	if t.root == nil {
		return *new(T)
	}

	var beginning = t.root
	for beginning.right != nil {
		beginning = beginning.right
	}

	if beginning == nil {
		for beginning.left != nil {
			beginning = beginning.left
		}
	}

	if beginning == nil {
		return *new(T)
	}

	return beginning.Value
}

// GetSize works in O(1) time, O(n) space
func (t *AvlTree[T]) GetSize() int {
	return t.size
}

////////////////////// AVL Tree Node //////////////////////

type AvlTreeNode[T comparable] struct {
	Value      T
	left       *AvlTreeNode[T]
	right      *AvlTreeNode[T]
	height     int
	comparator func(a, b T) int
}

func (n *AvlTreeNode[T]) init() *AvlTreeNode[T] {
	n.height = 1
	n.left = nil
	n.right = nil
	return n
}

func (n *AvlTreeNode[T]) get(val T) *AvlTreeNode[T] {
	var node *AvlTreeNode[T]

	compareResult := n.comparator(val, n.Value)
	if compareResult < 0 {
		if n.left != nil {
			node = n.left.get(val)
		}
	} else if compareResult > 0 {
		if n.right != nil {
			node = n.right.get(val)
		}
	} else {
		node = n
	}

	return node
}

func (n *AvlTreeNode[T]) rotateRight() *AvlTreeNode[T] {
	leftNode := n.left

	// Rotation
	leftNode.right, n.left = n, leftNode.right

	// Update heights
	n.height = n.getMaxHeight() + 1
	leftNode.height = leftNode.getMaxHeight() + 1

	return leftNode
}

func (n *AvlTreeNode[T]) rotateLeft() *AvlTreeNode[T] {
	rightNode := n.right

	// Rotation
	rightNode.left, n.right = n, rightNode.left

	// Update heights
	n.height = n.getMaxHeight() + 1
	rightNode.height = rightNode.getMaxHeight() + 1

	return rightNode
}

func (n *AvlTreeNode[T]) getMin() *AvlTreeNode[T] {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *AvlTreeNode[T]) getMaxHeight() int {
	rightHeight := getHeight(n.right)
	leftHeight := getHeight(n.left)

	if rightHeight > leftHeight {
		return rightHeight
	} else {
		return leftHeight
	}
}

////////////////////// Helper functions //////////////////////

func insertNode[T comparable](n *AvlTreeNode[T], value T, added *bool, comparator func(a, b T) int) *AvlTreeNode[T] {
	if n == nil {
		*added = true
		return (&AvlTreeNode[T]{
			Value:      value,
			comparator: comparator,
		}).init()
	}
	c := n.comparator(value, n.Value)
	if c > 0 {
		n.right = insertNode(n.right, value, added, comparator)
	} else if c < 0 {
		n.left = insertNode(n.left, value, added, comparator)
	} else {
		n.Value = value
		*added = false
		return n
	}

	n.height = n.getMaxHeight() + 1
	c = getChildrenHeightDiff[T](n)

	if c > 1 {
		c = n.left.comparator(value, n.left.Value)
		if c < 0 {
			return n.rotateRight()
		} else if c > 0 {
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	} else if c < -1 {
		c = n.right.comparator(value, n.right.Value)
		if c > 0 {
			return n.rotateLeft()
		} else if c < 0 {
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}
	return n
}

func deleteNode[T comparable](n *AvlTreeNode[T], value T, deleted *bool) *AvlTreeNode[T] {
	if n == nil {
		return n
	}

	c := n.comparator(value, n.Value)

	if c < 0 {
		n.left = deleteNode(n.left, value, deleted)
	} else if c > 0 {
		n.right = deleteNode(n.right, value, deleted)
	} else {
		if n.left == nil {
			t := n.right
			n.init()
			*deleted = true
			return t
		} else if n.right == nil {
			t := n.left
			n.init()
			*deleted = true
			return t
		} else {
			t := n.right.getMin()
			n.Value = t.Value
			n.right = deleteNode(n.right, t.Value, deleted)
			*deleted = true
		}
	}

	//re-getChildrenHeightDiff
	if n == nil {
		return n
	}
	n.height = n.getMaxHeight() + 1
	bal := getChildrenHeightDiff(n)
	if bal > 1 {
		if getChildrenHeightDiff(n.left) >= 0 {
			return n.rotateRight()
		}
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	} else if bal < -1 {
		if getChildrenHeightDiff(n.right) <= 0 {
			return n.rotateLeft()
		}
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func getHeight[T comparable](n *AvlTreeNode[T]) int {
	if n != nil {
		return n.height
	}

	return 0
}

func getChildrenHeightDiff[T comparable](n *AvlTreeNode[T]) int {
	if n == nil {
		return 0
	}

	return getHeight(n.left) - getHeight(n.right)
}
