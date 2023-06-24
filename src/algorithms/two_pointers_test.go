package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoPointersSimple(t *testing.T) {
	assert.Equal(t, []int{0, 1, 9, 25, 81}, TwoPointersSimple([]int{-3, -1, 0, 5, 9}))
	assert.Equal(t, []int{0, 1, 9, 25, 81}, TwoPointersSimple([]int{0, 1, 3, 5, 9}))
	assert.Equal(t, []int{0, 1, 9, 25, 81}, TwoPointersSimple([]int{-9, -5, -3, -1, 0}))
}
