package stdlib

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
