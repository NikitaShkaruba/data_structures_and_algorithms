package stdlib

import "math"

////////////////////// Math functions //////////////////////

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Abs[T Ordered](a T) T {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Pow[T Ordered](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}
