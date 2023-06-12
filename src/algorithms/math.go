package algorithms

////////////////////// Math functions //////////////////////

// ordered can be replaced with constraints.Ordered from "golang.org/x/exp/constraints".
// But because leetcode.com doesn't allow to use external libs, i've written mine
type ordered interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Max[T ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min[T ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Abs[T ordered](a T) T {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
