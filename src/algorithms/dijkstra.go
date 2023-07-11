package algorithms

import (
	"container/heap"
	"math"
)

//////////////////////// Dijkstra high-level ////////////////////////
// You can find this data structure overview in docs/algorithms/dijkstra.md

// DijkstraToTarget finds the shortest paths to every node in a sourceId connected component
func DijkstraToTarget(edges map[int]map[int]int, sourceId int, destinationId int) int {
	minPaths := dijkstra(edges, sourceId, destinationId)
	return minPaths[destinationId]
}

// DijkstraToEvery finds the shortest paths to every node in a sourceId connected component
func DijkstraToEvery(edges map[int]map[int]int, sourceId int) map[int]int {
	return dijkstra(edges, sourceId, unspecifiedTargetNodeId)
}

//////////////////////// Dijkstra low-level ////////////////////////

var unspecifiedTargetNodeId = math.MaxInt

type Edge struct {
	to   int
	cost int
}

// Time complexity: O(E * log(V))
// Space complexity: O(V)
func dijkstra(edges map[int]map[int]int, sourceId int, targetId int) map[int]int {
	minPaths := make(map[int]int)

	h := newDijkstraIndexedHeap()
	heap.Push(&h, Edge{
		to:   sourceId,
		cost: 0,
	})

	for len(h.values) != 0 {
		maxCostEdge := heap.Pop(&h).(Edge)
		minPaths[maxCostEdge.to] = maxCostEdge.cost

		// Only needed in DijkstraToDestination
		if targetId != unspecifiedTargetNodeId && maxCostEdge.to == targetId {
			return minPaths
		}

		for neighbourNode := range edges[maxCostEdge.to] {
			// We've already visited, so the new edge will not be the shortest -> don't include it
			if _, ok := minPaths[neighbourNode]; ok {
				continue
			}

			costOfGoingToNeighbour := edges[maxCostEdge.to][neighbourNode]
			newEdge := Edge{
				to:   neighbourNode,
				cost: maxCostEdge.cost + costOfGoingToNeighbour,
			}

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

//////////////////////// Dijkstra Indexed Heap ////////////////////////
// TODO: reuse heap from data_structures

type DijkstraIndexedHeap struct {
	values       []Edge
	valueIndexes map[int]int
}

func newDijkstraIndexedHeap() DijkstraIndexedHeap {
	return DijkstraIndexedHeap{
		values:       make([]Edge, 0),
		valueIndexes: make(map[int]int),
	}
}

func (h DijkstraIndexedHeap) Less(i, j int) bool {
	// Max heap
	return h.values[i].cost < h.values[j].cost
}
func (h *DijkstraIndexedHeap) Push(x interface{}) {
	val := x.(Edge)

	h.values = append(h.values, val)
	h.valueIndexes[val.to] = len(h.values) - 1
}
func (h *DijkstraIndexedHeap) Pop() interface{} {
	val := (h.values)[len(h.values)-1]

	h.values = h.values[:len(h.values)-1]
	delete(h.valueIndexes, val.to)

	return val
}
func (h DijkstraIndexedHeap) Swap(i, j int) {
	h.valueIndexes[h.values[i].to], h.valueIndexes[h.values[j].to] = h.valueIndexes[h.values[j].to], h.valueIndexes[h.values[i].to]
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
func (h DijkstraIndexedHeap) Len() int {
	return len(h.values)
}

func (h *DijkstraIndexedHeap) removeByNodeId(toNodeId int) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return
	}

	heap.Remove(h, i)
}
func (h *DijkstraIndexedHeap) updateByNodeId(toNodeId int, edge Edge) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return
	}

	heap.Remove(h, i)
	heap.Push(h, edge)
}
func (h *DijkstraIndexedHeap) getByNodeId(toNodeId int) (Edge, bool) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return Edge{}, false
	}

	return h.values[i], true
}
