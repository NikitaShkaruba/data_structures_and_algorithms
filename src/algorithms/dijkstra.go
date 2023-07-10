package algorithms

import (
	"container/heap"
	"math"
)

//////////////////////// Dijkstra high-level ////////////////////////
// You can find this data structure overview in docs/algorithms/dijkstra.md
// TODO: make weight a templated parameter
// TODO: add to leetcode_library.template

func DijkstraToDestination(edges map[int]map[int]float64, sourceId int, destinationId int) float64 {
	minPaths := dijkstra(edges, sourceId, destinationId)
	return minPaths[destinationId]
}

func DijkstraToEvery(edges map[int]map[int]float64, sourceId int) map[int]float64 {
	return dijkstra(edges, sourceId, prematureExitNotNeeded)
}

//////////////////////// Dijkstra low-level ////////////////////////

var prematureExitNotNeeded = math.MaxInt

type Edge struct {
	to   int
	cost float64
}

// Time complexity: O(E * log(V))
// Space complexity: O(V)
func dijkstra(edges map[int]map[int]float64, sourceId int, exitPrematurelyOnDestination int) map[int]float64 {
	minPaths := make(map[int]float64)

	h := newIndexedHeap()
	heap.Push(&h, Edge{
		to:   sourceId,
		cost: 1,
	})

	for len(h.values) != 0 {
		maxCostEdge := heap.Pop(&h).(Edge)
		minPaths[maxCostEdge.to] = maxCostEdge.cost

		// Only needed in DijkstraToDestination
		if exitPrematurelyOnDestination != prematureExitNotNeeded && maxCostEdge.to == exitPrematurelyOnDestination {
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
				cost: maxCostEdge.cost * costOfGoingToNeighbour,
			}

			if existingEdge, existsInHeap := h.getByNodeId(neighbourNode); existsInHeap {
				if newEdge.cost > existingEdge.cost {
					h.updateByNodeId(neighbourNode, newEdge)
				}
			} else {
				heap.Push(&h, newEdge)
			}
		}
	}

	return minPaths
}

//////////////////////// Indexed Heap Definition ////////////////////////
// TODO: reuse heap from data_structures

type IndexedHeap struct {
	values       []Edge
	valueIndexes map[int]int
}

//////////////////////// Indexed Heap Constructors ////////////////////////

func newIndexedHeap() IndexedHeap {
	return IndexedHeap{
		values:       make([]Edge, 0),
		valueIndexes: make(map[int]int),
	}
}

//////////////////////// Indexed Heap Methods ////////////////////////

func (h IndexedHeap) Less(i, j int) bool {
	// Max heap
	return h.values[i].cost > h.values[j].cost
}
func (h *IndexedHeap) Push(x interface{}) {
	val := x.(Edge)

	h.values = append(h.values, val)
	h.valueIndexes[val.to] = len(h.values) - 1
}
func (h *IndexedHeap) Pop() interface{} {
	val := (h.values)[len(h.values)-1]

	h.values = h.values[:len(h.values)-1]
	delete(h.valueIndexes, val.to)

	return val
}
func (h IndexedHeap) Swap(i, j int) {
	h.valueIndexes[h.values[i].to], h.valueIndexes[h.values[j].to] = h.valueIndexes[h.values[j].to], h.valueIndexes[h.values[i].to]
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
func (h IndexedHeap) Len() int {
	return len(h.values)
}

func (h *IndexedHeap) removeByNodeId(toNodeId int) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return
	}

	heap.Remove(h, i)
}
func (h *IndexedHeap) updateByNodeId(toNodeId int, edge Edge) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return
	}

	heap.Remove(h, i)
	heap.Push(h, edge)
}
func (h *IndexedHeap) getByNodeId(toNodeId int) (Edge, bool) {
	i, ok := h.valueIndexes[toNodeId]
	if !ok {
		return Edge{}, false
	}

	return h.values[i], true
}
