package data_structures

////////////////////// Union Find //////////////////////
// You can find this data structure overview in docs/data_structures/union_find.md

type UnionFind struct {
	sizeMap        []int
	parentMap      []int
	componentCount int
	nodeCount      int
}

// NewUnionFind works in O(n) time, O(n) space
func NewUnionFind(nodeCount int) *UnionFind {
	sizeMap := make([]int, nodeCount)
	parentMap := make([]int, nodeCount)
	for i := 0; i < nodeCount; i++ {
		parentMap[i] = i // Self parent
		sizeMap[i] = 1   // Each component has originally only 1 node
	}

	return &UnionFind{
		sizeMap:        sizeMap,
		parentMap:      parentMap,
		componentCount: nodeCount,
		nodeCount:      nodeCount,
	}
}

// Union unionizes components containing nodes p and q, works in O(1) time, O(1) space.
func (uf *UnionFind) Union(p, q int) bool {
	// These nodes are already in the same group!
	if uf.AreNodesConnected(p, q) {
		return false
	}

	pRoot := uf.Find(p)
	qRoot := uf.Find(q)

	// Merge smaller component into the larger one.
	if uf.sizeMap[pRoot] < uf.sizeMap[qRoot] {
		uf.sizeMap[qRoot] += uf.sizeMap[pRoot]
		uf.parentMap[pRoot] = qRoot
		uf.sizeMap[pRoot] = 0
	} else {
		uf.sizeMap[pRoot] += uf.sizeMap[qRoot]
		uf.parentMap[qRoot] = pRoot
		uf.sizeMap[qRoot] = 0
	}

	// Since the roots found are different we know that the
	// number of components has decreased by one
	uf.componentCount--

	return true
}

// Find finds which component p belongs to, works in O(1) time and O(1) space.
func (uf *UnionFind) Find(node int) int {
	// Find the root of the component
	root := node
	for root != uf.parentMap[root] {
		root = uf.parentMap[root]
	}

	// Compress the path leading back to the root.
	// Doing this operation is called "path compression"
	// and is what gives us amortized time complexity.
	for node != root {
		next := uf.parentMap[node]
		uf.parentMap[node] = root
		node = next
	}

	return root
}

// AreNodesConnected works in O(1) time, O(1) space
func (uf *UnionFind) AreNodesConnected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// GetComponentSize works in O(1) time, O(1) space
func (uf *UnionFind) GetComponentSize(p int) int {
	parentNode := uf.Find(p)
	return uf.sizeMap[parentNode]
}

// GetNodeCount works in O(1) time, O(1) space
func (uf *UnionFind) GetNodeCount() int {
	return uf.nodeCount
}

// GetComponentCount works in O(1) time, O(1) space
func (uf *UnionFind) GetComponentCount() int {
	return uf.componentCount
}
