package data_structures

////////////////////// Segment Tree //////////////////////

type SegmentTree[T ordered] struct {
	nums []T
	root *segment[T]
}

// NewSegmentTree works in O(n) time, O(logn) space
func NewSegmentTree[T ordered](nums []T) *SegmentTree[T] {
	return &SegmentTree[T]{
		nums: nums,
		root: newSegment(0, len(nums)-1, nums),
	}
}

// Update works in O(logn) time, O(1) space
func (t *SegmentTree[T]) Update(i int, val T) {
	diff := val - t.nums[i]

	t.nums[i] = val

	cur := t.root
	for cur != nil {
		cur.sum += diff

		var nextCur *segment[T]
		if cur.leftSegment != nil && cur.leftSegment.l <= i && i <= cur.leftSegment.r {
			nextCur = cur.leftSegment
		} else if cur.rightSegment != nil && cur.rightSegment.l <= i && i <= cur.rightSegment.r {
			nextCur = cur.rightSegment
		}

		cur = nextCur
	}
}

// GetRangeSum works in O(logn) time, O(logn) space
func (t *SegmentTree[T]) GetRangeSum(l int, r int) T {
	return t.root.sumRange(l, r)
}

// GetFullRangeSum works in O(logn) time, O(logn) space
func (t *SegmentTree[T]) GetFullRangeSum() T {
	return t.GetRangeSum(0, len(t.nums)-1)
}

////////////////////// Segment //////////////////////

type segment[T ordered] struct {
	l, r         int
	sum          T
	leftSegment  *segment[T]
	rightSegment *segment[T]
}

func newSegment[T ordered](l, r int, nums []T) *segment[T] {
	if l == r {
		return &segment[T]{
			l:   l,
			r:   r,
			sum: nums[l],
		}
	}

	m := l + (r-l)/2
	left := newSegment(l, m, nums)
	right := newSegment(m+1, r, nums)

	return &segment[T]{
		l:            l,
		r:            r,
		sum:          left.sum + right.sum,
		leftSegment:  left,
		rightSegment: right,
	}
}

func (s *segment[T]) sumRange(l, r int) T {
	// Disjoint
	if r < s.l || l > s.r {
		return *new(T)
	}

	// Covers
	if l <= s.l && s.r <= r {
		// Covers
		return s.sum
	}

	// Else
	return s.leftSegment.sumRange(l, r) + s.rightSegment.sumRange(l, r)
}

////////////////////// Ordered //////////////////////

// ordered can be replaced with constraints.Ordered from "golang.org/x/exp/constraints".
// But because leetcode.com doesn't allow to use external libs, i've written mine
// Also, I'm copypasting it everywhere for easier copy-pasting
type ordered interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
