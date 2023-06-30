package data_structures

import "src/src/stdlib"

////////////////////// Segment Tree //////////////////////
// More information is in docs/data_structures/segment_tree.md

type SegmentTree[T stdlib.Ordered] struct {
	nums []T
	root *segment[T]
}

// NewSegmentTree works in O(n) time, O(logn) space
func NewSegmentTree[T stdlib.Ordered](nums []T) *SegmentTree[T] {
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

type segment[T stdlib.Ordered] struct {
	l, r         int
	sum          T
	leftSegment  *segment[T]
	rightSegment *segment[T]
}

func newSegment[T stdlib.Ordered](l, r int, nums []T) *segment[T] {
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
