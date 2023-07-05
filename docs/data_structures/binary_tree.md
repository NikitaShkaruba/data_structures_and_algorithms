# Binary tree

Binary tree is basically a **directed graphs** with **single connected component** and **without cycles**

If every element on the left if lesser then the value of the current node and every element on the right is bigger then the value of the current node, this tree is **Binary Search Tree**

### Traversals

You can traverse binary trees via `Depth First Search` or `Breadth First Search`. 

Pre-order DFS traversal:
```go
func dfs(node *Node) {
	if node == nil {
		return
	}

	// Do some logic in this node
	dfs(node.left)
	dfs(node.right)
}
```


In-order DFS traversal (gives you a sorted array if you have a binary search (!) tree):
```go
func dfs(node *Node) {
	if node == nil {
		return
	}

	dfs(node.left)
	// Do some logic in this node
	dfs(node.right)
}
```


Post-order DFS traversal:
```go
func dfs(node *Node) {
	if node == nil {
		return
	}

	dfs(node.left)
	dfs(node.right)
	// Do some logic in this node
}
```

Level-order BFS traversal:

```go
func bfs(root *Node) []int {
	queue := data_structures.NewQueueFromArray([]*Node{root})
	level := 0

	for queue.GetSize() != 0 {
		nodesToProcess := queue.GetSize()

		for nodesToProcess != 0 {
			node := queue.Dequeue()
			nodesToProcess--

			// Do some logic in this node

			if node.Left != nil {
				queue = queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue = queue.Enqueue(node.Right)
			}
		}

		level++
	}
}
```
