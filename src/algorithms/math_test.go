package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrimeFactors(t *testing.T) {
	maxNumber := 10000

	assert.Equal(t, []int{2, 5}, GetPrimeFactors(20, maxNumber))
	assert.Equal(t, []int{2, 3}, GetPrimeFactors(36, maxNumber))
	assert.Equal(t, []int{2, 3, 47}, GetPrimeFactors(282, maxNumber))
	assert.Equal(t, []int{2, 5, 23}, GetPrimeFactors(460, maxNumber))
	assert.Equal(t, []int{2, 17}, GetPrimeFactors(544, maxNumber))
	assert.Equal(t, []int{2, 5, 31}, GetPrimeFactors(1240, maxNumber))
}
