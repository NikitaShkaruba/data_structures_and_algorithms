package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstraOnSmallGraphs(t *testing.T) {
	edges := map[int]map[int]float64{
		0: {
			1: 0.5,
			2: 0.2,
		},
		1: {
			0: 0.5,
			2: 0.5,
		},
		2: {
			0: 0.2,
			1: 0.5,
		},
	}

	shortestPaths := DijkstraToEvery(edges, 0)
	assert.Equal(t, 1.0, DijkstraToDestination(edges, 0, 0))
	assert.Equal(t, 0.5, shortestPaths[1])
	assert.Equal(t, 0.5, DijkstraToDestination(edges, 0, 1))
	assert.Equal(t, 0.25, shortestPaths[2])
	assert.Equal(t, 0.25, DijkstraToDestination(edges, 0, 2))

	shortestPaths = DijkstraToEvery(edges, 1)
	assert.Equal(t, 1.0, DijkstraToDestination(edges, 1, 1))
	assert.Equal(t, 0.5, shortestPaths[0])
	assert.Equal(t, 0.5, DijkstraToDestination(edges, 1, 0))
	assert.Equal(t, 0.5, shortestPaths[2])
	assert.Equal(t, 0.5, DijkstraToDestination(edges, 1, 2))

	shortestPaths = DijkstraToEvery(edges, 2)
	assert.Equal(t, 1.0, DijkstraToDestination(edges, 2, 2))
	assert.Equal(t, 0.25, shortestPaths[0])
	assert.Equal(t, 0.25, DijkstraToDestination(edges, 2, 0))
	assert.Equal(t, 0.5, shortestPaths[1])
	assert.Equal(t, 0.5, DijkstraToDestination(edges, 2, 1))
}
