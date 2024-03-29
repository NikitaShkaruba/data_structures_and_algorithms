package stdlib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 4, Max(1, 4))
	assert.Equal(t, 3, Max(-10, 3))
	assert.Equal(t, 15, Max(0, 15))
	assert.Equal(t, 2, Max(-2, 2))
	assert.Equal(t, -2, Max(-2, -10))
	assert.Equal(t, 1, Max(1, 1))
	assert.Equal(t, 123456789, Max(123456789, 123456788))
}

func TestMax_DifferentGenericTypes(t *testing.T) {
	assert.Equal(t, 2, Max(1, 2))
	assert.Equal(t, 2.0, Max(1.0, 2.0))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 4))
	assert.Equal(t, -10, Min(-10, 3))
	assert.Equal(t, 0, Min(0, 15))
	assert.Equal(t, -2, Min(-2, 2))
	assert.Equal(t, -10, Min(-2, -10))
	assert.Equal(t, 1, Min(1, 1))
	assert.Equal(t, 123456788, Min(123456789, 123456788))
}

func TestMin_DifferentGenericTypes(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1.0, Min(1.0, 2.0))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, 7, Abs(-7))
	assert.Equal(t, 0, Abs(0))
	assert.Equal(t, 123456789, Abs(-123456789))
}

func TestAbs_DifferentGenericTypes(t *testing.T) {
	assert.Equal(t, 1, Abs(-1))
	assert.Equal(t, 1.0, Abs(-1.0))
}

func TestPow(t *testing.T) {
	assert.Equal(t, 9, Pow(3, 2))
	assert.Equal(t, 256, Pow(2, 8))
	assert.Equal(t, 5.832000000000001, Pow(1.8, 3.0))
	assert.Equal(t, 13.064889045094622, Pow(3.4, 2.1))
}

func TestCeil(t *testing.T) {
	assert.Equal(t, 2, Ceil(1.1))
	assert.Equal(t, 1, Ceil(0.9))
	assert.Equal(t, 1, Ceil(0.5))
	assert.Equal(t, 1, Ceil(0.3))
	assert.Equal(t, 1, Ceil(0.1))
	assert.Equal(t, 0, Ceil(-0.1))
	assert.Equal(t, 0, Ceil(-0.3))
	assert.Equal(t, 0, Ceil(-0.5))
	assert.Equal(t, 0, Ceil(-0.9))
	assert.Equal(t, -1, Ceil(-1.1))
}
func TestFloor(t *testing.T) {
	assert.Equal(t, 1, Floor(1.1))
	assert.Equal(t, 0, Floor(0.9))
	assert.Equal(t, 0, Floor(0.5))
	assert.Equal(t, 0, Floor(0.3))
	assert.Equal(t, 0, Floor(0.1))
	assert.Equal(t, -1, Floor(-0.1))
	assert.Equal(t, -1, Floor(-0.3))
	assert.Equal(t, -1, Floor(-0.5))
	assert.Equal(t, -1, Floor(-0.9))
	assert.Equal(t, -2, Floor(-1.1))
}
