package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var smallUndirectedGraphEdges = map[int]map[int]int{
	0: {
		1: 1,
		2: 6,
	},
	1: {
		0: 1,
		2: 2,
	},
	2: {
		0: 6,
		1: 2,
		3: 3,
	},
	3: {
		2: 3,
	},
}

func TestDijkstraOnSmallUndirectedGraph(t *testing.T) {
	edges := smallUndirectedGraphEdges

	shortestPaths := DijkstraToEvery(edges, 0)
	assert.Equal(t, 0, shortestPaths[0])
	assert.Equal(t, 1, shortestPaths[1])
	assert.Equal(t, 3, shortestPaths[2])
	assert.Equal(t, 6, shortestPaths[3])

	shortestPaths = DijkstraToEvery(edges, 1)
	assert.Equal(t, 1, shortestPaths[0])
	assert.Equal(t, 0, shortestPaths[1])
	assert.Equal(t, 2, shortestPaths[2])
	assert.Equal(t, 5, shortestPaths[3])

	shortestPaths = DijkstraToEvery(edges, 2)
	assert.Equal(t, 3, shortestPaths[0])
	assert.Equal(t, 2, shortestPaths[1])
	assert.Equal(t, 0, shortestPaths[2])
	assert.Equal(t, 3, shortestPaths[3])

	shortestPaths = DijkstraToEvery(edges, 3)
	assert.Equal(t, 6, shortestPaths[0])
	assert.Equal(t, 5, shortestPaths[1])
	assert.Equal(t, 3, shortestPaths[2])
	assert.Equal(t, 0, shortestPaths[3])
}

func TestDijkstraOnBigDirectedGraph(t *testing.T) {
	// https://www.tutorialspoint.com/design_and_analysis_of_algorithms/design_and_analysis_of_algorithms_shortest_paths.htm
	edges := map[int]map[int]int{
		0: {1: 5, 2: 2},
		1: {3: 3, 4: 7},
		2: {1: 2, 5: 9},
		3: {2: -2, 4: 2, 5: 6},
		4: {5: 5, 6: 8, 7: 7},
		5: {7: 2},
		6: {8: 4},
		7: {6: 3},
	}

	shortestPaths := DijkstraToEvery(edges, 0)
	assert.Equal(t, 0, shortestPaths[0])
	assert.Equal(t, 4, shortestPaths[1])
	assert.Equal(t, 2, shortestPaths[2])
	assert.Equal(t, 7, shortestPaths[3])
	assert.Equal(t, 9, shortestPaths[4])
	assert.Equal(t, 11, shortestPaths[5])
	assert.Equal(t, 16, shortestPaths[6])
	assert.Equal(t, 13, shortestPaths[7])
	assert.Equal(t, 20, shortestPaths[8])
}

func TestDijkstraToTargetOnSmallUndirectedGraph(t *testing.T) {
	edges := smallUndirectedGraphEdges

	shortestPaths := DijkstraToEvery(edges, 0)
	assert.Equal(t, 0, shortestPaths[0])
	assert.Equal(t, 0, DijkstraToTarget(edges, 0, 0))
	assert.Equal(t, 1, shortestPaths[1])
	assert.Equal(t, 1, DijkstraToTarget(edges, 0, 1))
	assert.Equal(t, 3, shortestPaths[2])
	assert.Equal(t, 3, DijkstraToTarget(edges, 0, 2))
	assert.Equal(t, 6, shortestPaths[3])
	assert.Equal(t, 6, DijkstraToTarget(edges, 0, 3))
}
