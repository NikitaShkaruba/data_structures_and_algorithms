# Binary tree

Binary tree is basically a **directed graphs** with a **single connected component** and **without cycles**

If every element on the left if lesser then the value of the current node and every element on the right is bigger then the value of the current node, this tree is **Binary Search Tree**

Source code: [src/data_structures/binary_tree.go](../../src/data_structures/binary_tree.go)

### Traversals

You can traverse binary trees via `Depth First Search` or `Breadth First Search`. 

Pre-order DFS traversal:

```go
func dfs[T any](node *data_structures.BinaryTreeNode[T]) {
	if node == nil {
		return
	}

	// Do some logic in this node
	dfs(node.Left)
	dfs(node.Right)
}
```

In-order DFS traversal (gives you a sorted array if you have a binary search (!) tree):

```go
func dfs[T any](node *data_structures.BinaryTreeNode[T]) {
	if node == nil {
		return
	}

	dfs(node.Left)
	// Do some logic in this node
	dfs(node.Right)
}
```

Post-order DFS traversal:

```go
func dfs[T any](node *data_structures.BinaryTreeNode[T]) {
	if node == nil {
		return
	}

	dfs(node.Left)
	dfs(node.Right)
	// Do some logic in this node
}
```

Level-order BFS traversal:

```go
func bfs[T any](root *data_structures.BinaryTreeNode[T]) {
	queue := data_structures.NewQueueFromArray(
		[]*data_structures.BinaryTreeNode[T]{root},
	)
	level := 0

	for queue.GetSize() != 0 {
		nodesToProcess := queue.GetSize()

		for nodesToProcess != 0 {
			node := queue.Dequeue()
			nodesToProcess--

			// Do some logic in this node

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		level++
	}
}

```

### Binary search

If your tree is sorted, you can use [binary search](../algorithms/binary_search.md) to quickly find if a value is
contained in a tree for `O(logn)` time.

```go
func binarySearch(node *data_structures.BinaryTreeNode[int], target int) bool {
	if node == nil {
		return false
	}

	if node.Val == target {
		return true
	}

	if target < node.Val {
		return binarySearch(node.Left, target)
	} else {
		return binarySearch(node.Right, target)
	}
}
```
