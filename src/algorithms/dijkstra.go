package algorithms

import (
	"math"
	"src/src/data_structures"
)

//////////////////////// Dijkstra high-level ////////////////////////
// You can find this algorithm overview in docs/algorithms/dijkstra.md

// DijkstraToTarget finds the shortest path from sourceId to targetId in a sourceId connected component
func DijkstraToTarget(edges map[int]map[int]int, sourceId int, targetId int) int {
	minPaths := dijkstra(edges, sourceId, targetId)
	return minPaths[targetId]
}

// DijkstraToEvery finds the shortest path from sourceId to every node in a sourceId connected component
func DijkstraToEvery(edges map[int]map[int]int, sourceId int) map[int]int {
	return dijkstra(edges, sourceId, dijkstraUnspecifiedTargetId)
}

//////////////////////// Dijkstra low-level ////////////////////////

// Time complexity: O(E * log(V))
// Space complexity: O(V)
func dijkstra(edges map[int]map[int]int, sourceId int, targetId int) map[int]int {
	minPaths := make(map[int]int)

	h := data_structures.NewEmptyIndexedHeap(func(a, b dijkstraEdge) bool {
		return a.cost < b.cost
	})
	h.Push(dijkstraEdge{
		to:   sourceId,
		cost: 0,
	})

	for h.GetSize() != 0 {
		minCostEdge := h.Pop()

		// We might already have a shortest path
		if _, ok := minPaths[minCostEdge.to]; ok {
			continue
		}
		minPaths[minCostEdge.to] = minCostEdge.cost

		// Only needed in DijkstraToDestination
		if targetId != dijkstraUnspecifiedTargetId && minCostEdge.to == targetId {
			return minPaths
		}

		for neighbourNode := range edges[minCostEdge.to] {
			// We've already visited, so the new edge will not be the shortest -> don't include it
			if _, ok := minPaths[neighbourNode]; ok {
				continue
			}

			costOfGoingToNeighbour := edges[minCostEdge.to][neighbourNode]
			h.Push(dijkstraEdge{
				to:   neighbourNode,
				cost: minCostEdge.cost + costOfGoingToNeighbour,
			})
		}
	}

	return minPaths
}

//////////////////////// Dijkstra helpers ////////////////////////

var dijkstraUnspecifiedTargetId = math.MaxInt

type dijkstraEdge struct {
	to   int
	cost int
}
