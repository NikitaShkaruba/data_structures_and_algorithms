package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFind_New(t *testing.T) {
	nodeCount := 10

	uf := NewUnionFind(nodeCount)

	assert.Equal(t, nodeCount, uf.GetComponentCount())
	assert.Equal(t, nodeCount, uf.GetNodeCount())

	for i := 0; i < nodeCount; i++ {
		assert.Equal(t, 1, uf.GetComponentSize(i))

		for j := i + 1; j < nodeCount; j++ {
			assert.False(t, uf.AreNodesConnected(i, j))
		}
	}
}

func TestUnionFind_UnionAreNodesConnected(t *testing.T) {
	uf := NewUnionFind(3)

	firstNode := 0
	secondNode := 1

	assert.False(t, uf.AreNodesConnected(firstNode, secondNode))
	assert.True(t, uf.Union(firstNode, secondNode))
	assert.True(t, uf.AreNodesConnected(firstNode, secondNode))

	thirdNode := 2

	assert.False(t, uf.AreNodesConnected(firstNode, thirdNode))
	assert.False(t, uf.AreNodesConnected(secondNode, thirdNode))
	assert.True(t, uf.Union(firstNode, thirdNode))
	assert.True(t, uf.AreNodesConnected(firstNode, thirdNode))
	assert.True(t, uf.AreNodesConnected(secondNode, thirdNode))

	assert.False(t, uf.Union(firstNode, secondNode))
	assert.False(t, uf.Union(firstNode, thirdNode))
}

func TestUnionFind_UnionGetComponentSize(t *testing.T) {
	uf := NewUnionFind(3)

	firstNode := 0
	secondNode := 1

	assert.Equal(t, 1, uf.GetComponentSize(firstNode))
	assert.Equal(t, 1, uf.GetComponentSize(secondNode))
	uf.Union(firstNode, secondNode)
	assert.Equal(t, 2, uf.GetComponentSize(firstNode))
	assert.Equal(t, 2, uf.GetComponentSize(secondNode))

	thirdNode := 2

	assert.Equal(t, 1, uf.GetComponentSize(thirdNode))
	uf.Union(firstNode, thirdNode)
	assert.Equal(t, firstNode, uf.Find(thirdNode))
	assert.Equal(t, 3, uf.GetComponentSize(firstNode))
	assert.Equal(t, 3, uf.GetComponentSize(secondNode))
	assert.Equal(t, 3, uf.GetComponentSize(thirdNode))
}

func TestUnionFind_UnionGetComponentCount(t *testing.T) {
	uf := NewUnionFind(3)

	firstNode := 0
	secondNode := 1
	thirdNode := 2

	assert.Equal(t, 3, uf.GetComponentCount())

	uf.Union(firstNode, secondNode)
	assert.Equal(t, 2, uf.GetComponentCount())

	uf.Union(firstNode, thirdNode)
	assert.Equal(t, 1, uf.GetComponentCount())
}

func TestUnionFind_Find(t *testing.T) {
	uf := NewUnionFind(3)

	firstNode := 0
	secondNode := 1

	assert.Equal(t, firstNode, uf.Find(firstNode))
	assert.Equal(t, secondNode, uf.Find(secondNode))
	uf.Union(firstNode, secondNode)
	assert.Equal(t, firstNode, uf.Find(firstNode))
	assert.Equal(t, firstNode, uf.Find(secondNode))

	thirdNode := 2

	assert.Equal(t, thirdNode, uf.Find(thirdNode))
	uf.Union(firstNode, thirdNode)
	assert.Equal(t, firstNode, uf.Find(thirdNode))
}
