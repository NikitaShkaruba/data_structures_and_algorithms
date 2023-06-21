package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitchSegmentTree_GetRangeSum(t *testing.T) {
	switches := []int{1, 0, 0, 1, 1, 0}
	st := NewSwitchSegmentTree(switches)

	// Ranges of length 1
	for i := range switches {
		assert.Equal(t, switches[i], st.GetRangeSum(i, i))
	}

	// Ranges of length 2
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 0, st.GetRangeSum(1, 2))
	assert.Equal(t, 1, st.GetRangeSum(2, 3))
	assert.Equal(t, 2, st.GetRangeSum(3, 4))
	assert.Equal(t, 1, st.GetRangeSum(4, 5))

	// Ranges of length 3
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(1, 3))
	assert.Equal(t, 2, st.GetRangeSum(2, 4))
	assert.Equal(t, 2, st.GetRangeSum(3, 5))

	// Ranges of length 4
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 2, st.GetRangeSum(1, 4))
	assert.Equal(t, 2, st.GetRangeSum(2, 5))

	// Ranges of length 5
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(1, 5))

	// Ranges of length 6
	assert.Equal(t, 3, st.GetRangeSum(0, 5))
}

func TestSwitchSegmentTree_GetFullRangeSum(t *testing.T) {
	switches := []int{1, 0, 1, 1, 0}
	st := NewSwitchSegmentTree(switches)
	assert.Equal(t, 3, st.GetFullRangeSum())

	switches = []int{0, 0, 0}
	st = NewSwitchSegmentTree(switches)
	assert.Equal(t, 0, st.GetFullRangeSum())

	switches = []int{1, 1, 1, 1}
	st = NewSwitchSegmentTree(switches)
	assert.Equal(t, 4, st.GetFullRangeSum())
}

func TestSwitchSegmentTree_Flip(t *testing.T) {
	switches := []int{1, 0, 0, 1, 1, 0}
	st := NewSwitchSegmentTree(switches)

	assert.Equal(t, 3, st.GetRangeSum(0, 5))
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 1, st.GetRangeSum(0, 0))

	st.Flip(0)
	assert.Equal(t, 2, st.GetRangeSum(0, 5))
	assert.Equal(t, 2, st.GetRangeSum(0, 4))
	assert.Equal(t, 1, st.GetRangeSum(0, 3))
	assert.Equal(t, 0, st.GetRangeSum(0, 2))
	assert.Equal(t, 0, st.GetRangeSum(0, 1))
	assert.Equal(t, 0, st.GetRangeSum(0, 0))

	st.Flip(0)
	assert.Equal(t, 3, st.GetRangeSum(0, 5))
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 1, st.GetRangeSum(0, 0))
}

func TestSwitchSegmentTree_FlipRange(t *testing.T) {
	switches := []int{1, 0, 0, 1, 1, 0}
	st := NewSwitchSegmentTree(switches)

	assert.Equal(t, 3, st.GetRangeSum(0, 5))
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 1, st.GetRangeSum(0, 0))

	st.FlipRange(0, 1)
	assert.Equal(t, 3, st.GetRangeSum(0, 5))
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 0, st.GetRangeSum(0, 0))
	st.FlipRange(0, 1)
	assert.Equal(t, 3, st.GetRangeSum(0, 5))
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 1, st.GetRangeSum(0, 0))

	st.FlipRange(1, 2)
	assert.Equal(t, 5, st.GetRangeSum(0, 5))
	assert.Equal(t, 5, st.GetRangeSum(0, 4))
	assert.Equal(t, 4, st.GetRangeSum(0, 3))
	assert.Equal(t, 3, st.GetRangeSum(0, 2))
	assert.Equal(t, 2, st.GetRangeSum(0, 1))
	assert.Equal(t, 1, st.GetRangeSum(0, 0))
	st.FlipRange(1, 2)
	assert.Equal(t, 3, st.GetRangeSum(0, 5))
	assert.Equal(t, 3, st.GetRangeSum(0, 4))
	assert.Equal(t, 2, st.GetRangeSum(0, 3))
	assert.Equal(t, 1, st.GetRangeSum(0, 2))
	assert.Equal(t, 1, st.GetRangeSum(0, 1))
	assert.Equal(t, 1, st.GetRangeSum(0, 0))
}

func TestSwitchSegmentTree_IsActive(t *testing.T) {
	switches := []int{1, 0, 0, 1, 1, 0}
	st := NewSwitchSegmentTree(switches)

	assert.True(t, st.IsActive(0))
	assert.False(t, st.IsActive(1))
	assert.False(t, st.IsActive(2))
	assert.True(t, st.IsActive(3))
	assert.True(t, st.IsActive(4))
	assert.False(t, st.IsActive(5))
}
