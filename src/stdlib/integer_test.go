package stdlib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntegerLength(t *testing.T) {
	assert.Equal(t, 5, GetIntegerLength(-12345))
	assert.Equal(t, 4, GetIntegerLength(-1234))
	assert.Equal(t, 3, GetIntegerLength(-123))
	assert.Equal(t, 2, GetIntegerLength(-12))
	assert.Equal(t, 1, GetIntegerLength(-1))
	assert.Equal(t, 1, GetIntegerLength(0))
	assert.Equal(t, 1, GetIntegerLength(1))
	assert.Equal(t, 2, GetIntegerLength(12))
	assert.Equal(t, 3, GetIntegerLength(123))
	assert.Equal(t, 4, GetIntegerLength(1234))
	assert.Equal(t, 5, GetIntegerLength(12345))
}

func TestGetIntegerDigit(t *testing.T) {
	assert.Equal(t, 5, GetIntegerDigit(-12345, 0))
	assert.Equal(t, 4, GetIntegerDigit(-12345, 1))
	assert.Equal(t, 3, GetIntegerDigit(-12345, 2))
	assert.Equal(t, 2, GetIntegerDigit(-12345, 3))
	assert.Equal(t, 1, GetIntegerDigit(-12345, 4))

	assert.Equal(t, 5, GetIntegerDigit(12345, 0))
	assert.Equal(t, 4, GetIntegerDigit(12345, 1))
	assert.Equal(t, 3, GetIntegerDigit(12345, 2))
	assert.Equal(t, 2, GetIntegerDigit(12345, 3))
	assert.Equal(t, 1, GetIntegerDigit(12345, 4))
}

func TestPow(t *testing.T) {
	assert.Equal(t, 9, Pow(3, 2))
	assert.Equal(t, 256, Pow(2, 8))
	assert.Equal(t, 5.832000000000001, Pow(1.8, 3.0))
	assert.Equal(t, 13.064889045094622, Pow(3.4, 2.1))
}
