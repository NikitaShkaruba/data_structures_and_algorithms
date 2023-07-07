# Graphs

- Can be **directed** / **undirected**
- Can have multiple **connected components**
- The amount of edges entering a node is its **indegree,** the amount of edges exiting a node is its **outdegree**
- Nodes that are connected by an edge are called **neighbours**

### Graph construction

Graph can be given in different ways
1. **array of edges** of size m `edges = [[0, 1], [1, 2], [2, 0], [2, 3]]`, finding neighbours is O(n^2), which is a lot, we can preprocess it to a map.
2. **adjacency matrix** of size n^2 ****`connections = [[0,0,1], [1,0,1], [1,0,0]]`, finding neighbours is O(n), which is a lot, we can preprocess it to a map.
3. **adjacency list** of size n `graph = [[1], [2], [0, 3], []]`, finding neighbours is O(1), which is very good, but we store empty arrays, which is not ideal. Preprocessing to a map is not needed, this is the most convenient format.
4. **matrix** of size n `vertexes = [[0,0,1], [1,0,1], [1,0,0]]`, in this input 1 is a vertex, 0 is not and there are max only 4 neighbours at `(row - 1, col), (row, col - 1), (row + 1, col), (row, col + 1)`. Finding neighbours is O(1) which is good, no preprocessing needed.

### Depth First Search (DFS)

80% of graph problems are solved with DFS with some conditions. So you'd better get acquainted with it.

```go
package s

// dfs works in O(n + e) time, O(n) space
func dfs(v int, adjacencyList [][]int, seen map[int]struct{}) {
	if _, ok := seen[v]; ok {
		return
	}
	seen[v] = struct{}{}

	for i := range adjacencyList[v] {
		dfs(adjacencyList[v][i], adjacencyList, seen)
	}
}

```

Finding amount of connected components with DFS.

```go
package s

func findConnectedComponentsCount(adjacencyList [][]int) int {
	connectedComponentsCount := 0
	seen := make(map[int]struct{})

	for i := range adjacencyList {
		if _, ok := seen[i]; ok {
			// It already belongs to some connected component
			continue
		}

		// We are starting a new connected component
		connectedComponentsCount++
		dfs(i, adjacencyList, seen)
	}

	return connectedComponentsCount
}

// Mark all the parts of our connected component
func dfs(v int, adjacencyList [][]int, seen map[int]struct{}) {
	if _, ok := seen[v]; ok {
		return
	}
	seen[v] = struct{}{}

	for i := range adjacencyList[v] {
		dfs(adjacencyList[v][i], adjacencyList, seen)
	}
}
```

It's even easier to do with [Union-Find](../data_structures/union_find.md) by the way.

### Breadth First Search (BFS)

- Can find the shortest path possible
- Harder to implement than DFS, but finds a path too, so use it only when necessary

### Tree-graphs

- If a graph has `n` nodes and `n-1` edges, it’s a tree with possibly more than 2 children.
- There will be no cycles
- E.g. https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/

### Bipartite graph

A bipartite graph (or bigraph) is a graph whose vertices can be divided into two disjoint and independent sets U and V
where there aren't any edges between nodes of the same set. It can be solved with Union Find.

```go
// DislikesArray contains pairs that need to be at the different sets
func isBipartitionPossible(n int, dislikesArray [][2]int) bool {
	dislikes := make(map[int][]int)
	for _, d := range dislikesArray {
		dislikes[d[0]-1] = append(dislikes[d[0]-1], d[1]-1)
		dislikes[d[1]-1] = append(dislikes[d[1]-1], d[0]-1)
	}

	uf := data_structures.NewUnionFind(n)

	for v := 0; v < n; v++ {
		if len(dislikes[v]) == 0 {
			continue
		}

		for i := 1; i < len(dislikes[v]); i++ {
			if uf.Find(v) == uf.Find(dislikes[v][i]) {
				return false
			}

			uf.Union(dislikes[v][0], dislikes[v][i])
		}
	}

	return true
}
```
