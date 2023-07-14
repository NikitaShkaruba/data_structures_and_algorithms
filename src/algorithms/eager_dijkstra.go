package algorithms

import (
	"container/heap"
)

//////////////////////// Eager dijkstra high-level ////////////////////////
// You can find this algorithm overview in docs/algorithms/dijkstra.md

// EagerDijkstraToTarget finds the shortest path from sourceId to targetId in a sourceId connected component
func EagerDijkstraToTarget(edges map[int]map[int]int, sourceId int, targetId int) int {
	minPaths := eagerDijkstra(edges, sourceId, targetId)
	return minPaths[targetId]
}

// EagerDijkstraToEvery finds the shortest path from sourceId to every node in a sourceId connected component
func EagerDijkstraToEvery(edges map[int]map[int]int, sourceId int) map[int]int {
	return eagerDijkstra(edges, sourceId, dijkstraUnspecifiedTargetId)
}

//////////////////////// Eager dijkstra low-level ////////////////////////

// Time complexity: O(E * log(V))
// Space complexity: O(V)
func eagerDijkstra(edges map[int]map[int]int, sourceId int, targetId int) map[int]int {
	minPaths := make(map[int]int)

	h := newEagerDijkstraIndexedHeap()
	heap.Push(&h, dijkstraEdge{
		to:   sourceId,
		cost: 0,
	})

	for len(h.values) != 0 {
		minCostEdge := heap.Pop(&h).(dijkstraEdge)
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
			newEdge := dijkstraEdge{
				to:   neighbourNode,
				cost: minCostEdge.cost + costOfGoingToNeighbour,
			}

			// Try to replace existing edge if your path is faster.
			if existingEdge, existsInHeap := h.getByNodeId(neighbourNode); existsInHeap {
				if newEdge.cost < existingEdge.cost {
					h.updateByNodeId(neighbourNode, newEdge)
				}
			} else {
				heap.Push(&h, newEdge)
			}
		}
	}

	return minPaths
}

//////////////////////// Eager dijkstra Indexed Heap ////////////////////////

// EagerDijkstraIndexedHeap heap is a modified version of indexed heap that stores
// edges in values, and allows to access them by nodeId
type EagerDijkstraIndexedHeap struct {
	values       []dijkstraEdge
	valueIndexes map[int]int
}

func newEagerDijkstraIndexedHeap() EagerDijkstraIndexedHeap {
	return EagerDijkstraIndexedHeap{
		values:       make([]dijkstraEdge, 0),
		valueIndexes: make(map[int]int),
	}
}

func (h EagerDijkstraIndexedHeap) Less(i, j int) bool {
	// Min heap
	return h.values[i].cost < h.values[j].cost
}

func (h *EagerDijkstraIndexedHeap) Push(x interface{}) {
	val := x.(dijkstraEdge)

	h.values = append(h.values, val)
	h.valueIndexes[val.to] = len(h.values) - 1
}

func (h *EagerDijkstraIndexedHeap) Pop() interface{} {
	val := (h.values)[len(h.values)-1]

	h.values = h.values[:len(h.values)-1]
	delete(h.valueIndexes, val.to)

	return val
}

func (h EagerDijkstraIndexedHeap) Swap(i, j int) {
	h.valueIndexes[h.values[i].to], h.valueIndexes[h.values[j].to] = h.valueIndexes[h.values[j].to], h.valueIndexes[h.values[i].to]
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h EagerDijkstraIndexedHeap) Len() int {
	return len(h.values)
}

func (h *EagerDijkstraIndexedHeap) removeByNodeId(toNodeId int) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return
	}

	heap.Remove(h, i)
}

func (h *EagerDijkstraIndexedHeap) updateByNodeId(toNodeId int, edge dijkstraEdge) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return
	}

	heap.Remove(h, i)
	heap.Push(h, edge)
}

func (h *EagerDijkstraIndexedHeap) getByNodeId(toNodeId int) (dijkstraEdge, bool) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return dijkstraEdge{}, false
	}

	return h.values[i], true
}
