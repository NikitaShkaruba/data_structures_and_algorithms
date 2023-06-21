package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentTree_GetFullRangeSum(t *testing.T) {
	st := NewSegmentTree([]int{10, 5, 4, -3, 2})
	assert.Equal(t, 18, st.GetFullRangeSum())

	st = NewSegmentTree([]int{-1, -2, -3})
	assert.Equal(t, -6, st.GetFullRangeSum())

	st = NewSegmentTree([]int{0, 0, 0})
	assert.Equal(t, 0, st.GetFullRangeSum())
}

func TestSegmentTree_GetRangeSum(t *testing.T) {
	st := NewSegmentTree([]int{10, 5, 4, -3, 2})

	assert.Equal(t, 10, st.GetRangeSum(0, 0))
	assert.Equal(t, 15, st.GetRangeSum(0, 1))
	assert.Equal(t, 19, st.GetRangeSum(0, 2))
	assert.Equal(t, 16, st.GetRangeSum(0, 3))
	assert.Equal(t, 18, st.GetRangeSum(0, 4))

	assert.Equal(t, 5, st.GetRangeSum(1, 1))
	assert.Equal(t, 9, st.GetRangeSum(1, 2))
	assert.Equal(t, 6, st.GetRangeSum(1, 3))
	assert.Equal(t, 8, st.GetRangeSum(1, 4))

	assert.Equal(t, 4, st.GetRangeSum(2, 2))
	assert.Equal(t, 1, st.GetRangeSum(2, 3))
	assert.Equal(t, 3, st.GetRangeSum(2, 4))

	assert.Equal(t, -3, st.GetRangeSum(3, 3))
	assert.Equal(t, -1, st.GetRangeSum(3, 4))

	assert.Equal(t, 2, st.GetRangeSum(4, 4))
}

func TestSegmentTree_Update(t *testing.T) {
	st := NewSegmentTree([]int{10, 5, 4, -3, 2})

	assert.Equal(t, 18, st.GetFullRangeSum())
	assert.Equal(t, 4, st.GetRangeSum(2, 2))
	assert.Equal(t, 9, st.GetRangeSum(1, 2))
	assert.Equal(t, 1, st.GetRangeSum(2, 3))
	assert.Equal(t, 6, st.GetRangeSum(1, 3))

	st.Update(2, 2)
	assert.Equal(t, 16, st.GetFullRangeSum())
	assert.Equal(t, 2, st.GetRangeSum(2, 2))
	assert.Equal(t, 7, st.GetRangeSum(1, 2))
	assert.Equal(t, -1, st.GetRangeSum(2, 3))
	assert.Equal(t, 4, st.GetRangeSum(1, 3))

	st.Update(2, 0)
	assert.Equal(t, 14, st.GetFullRangeSum())
	assert.Equal(t, 0, st.GetRangeSum(2, 2))
	assert.Equal(t, 5, st.GetRangeSum(1, 2))
	assert.Equal(t, -3, st.GetRangeSum(2, 3))
	assert.Equal(t, 2, st.GetRangeSum(1, 3))

	st.Update(2, 4)
	assert.Equal(t, 18, st.GetFullRangeSum())
	assert.Equal(t, 4, st.GetRangeSum(2, 2))
	assert.Equal(t, 9, st.GetRangeSum(1, 2))
	assert.Equal(t, 1, st.GetRangeSum(2, 3))
	assert.Equal(t, 6, st.GetRangeSum(1, 3))
}

func TestSegmentTree_TestDifferentGenericTypes(t *testing.T) {
	intSegmentTree := NewSegmentTree([]int{1, 2, 3})
	assert.Equal(t, 6, intSegmentTree.GetFullRangeSum())

	floatSegmentTree := NewSegmentTree([]float64{1.0, 2.0, 3.0})
	assert.Equal(t, 6.0, floatSegmentTree.GetFullRangeSum())
}
