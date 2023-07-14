package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 4, GCD(8, 12))
	assert.Equal(t, 6, GCD(54, 24))
	assert.Equal(t, 14, GCD(42, 56))
	assert.Equal(t, 1, GCD(13, 48))
}

func TestGetPrimeFactors(t *testing.T) {
	maxNumber := 10000

	assert.Equal(t, []int{2, 5}, GetPrimeFactors(20, maxNumber))
	assert.Equal(t, []int{2, 3}, GetPrimeFactors(36, maxNumber))
	assert.Equal(t, []int{2, 3, 47}, GetPrimeFactors(282, maxNumber))
	assert.Equal(t, []int{2, 5, 23}, GetPrimeFactors(460, maxNumber))
	assert.Equal(t, []int{2, 17}, GetPrimeFactors(544, maxNumber))
	assert.Equal(t, []int{2, 5, 31}, GetPrimeFactors(1240, maxNumber))
}
