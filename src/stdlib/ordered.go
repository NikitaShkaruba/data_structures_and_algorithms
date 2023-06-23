package stdlib

////////////////////// Ordered //////////////////////

// Ordered can be replaced with constraints.Ordered from "golang.org/x/exp/constraints".
// But because leetcode.com doesn't allow to use external libs, so it was copypasted
type Ordered interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
