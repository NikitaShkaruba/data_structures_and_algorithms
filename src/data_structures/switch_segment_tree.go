package data_structures

////////////////////// Switch Segment Tree //////////////////////
// More information is in docs/data_structures/segment_tree.md

// SwitchSegmentTree is a data structure that contain array of switches (1 or 0).
// It can flip any switch range quickly for O(logn) time instead of usual O(n) time for arrays
type SwitchSegmentTree struct {
	root *switchSegment
}

// NewSwitchSegmentTree works in O(n) time, O(logn) space
func NewSwitchSegmentTree(switches []int) *SwitchSegmentTree {
	return &SwitchSegmentTree{
		root: newSwitchSegment(0, len(switches)-1, switches),
	}
}

// Flip works in O(logn) time, O(logn) space
func (t *SwitchSegmentTree) Flip(i int) {
	t.FlipRange(i, i)
}

// FlipRange works in O(logn) time, O(logn) space
func (t *SwitchSegmentTree) FlipRange(l, r int) {
	t.root.flipRange(l, r)
}

// GetRangeSum works in O(logn) time, O(logn) space
func (t *SwitchSegmentTree) GetRangeSum(l int, r int) int {
	return t.root.sumRange(l, r)
}

// GetFullRangeSum works in O(logn) time, O(logn) space
func (t *SwitchSegmentTree) GetFullRangeSum() int {
	return t.root.sumRange(t.root.l, t.root.r)
}

// IsActive works in O(logn) time, O(logn) space
func (t *SwitchSegmentTree) IsActive(i int) bool {
	return t.GetRangeSum(i, i) == 1
}

////////////////////// Switch Segment //////////////////////

type switchSegment struct {
	// l and r are the index bounds of a segment
	l, r         int
	sum, count   int
	leftSegment  *switchSegment
	rightSegment *switchSegment
	flipped      bool
}

func newSwitchSegment(l, r int, switches []int) *switchSegment {
	if l == r {
		return &switchSegment{
			l:     l,
			r:     r,
			sum:   switches[l],
			count: 1,
		}
	}

	m := l + (r-l)/2
	left := newSwitchSegment(l, m, switches)
	right := newSwitchSegment(m+1, r, switches)

	return &switchSegment{
		l:            l,
		r:            r,
		sum:          left.sum + right.sum,
		count:        left.count + right.count,
		leftSegment:  left,
		rightSegment: right,
	}
}

func (s *switchSegment) sumRange(l, r int) int {
	// Disjoint
	if r < s.l || l > s.r {
		return 0
	}

	// Covers
	if l <= s.l && s.r <= r {
		return s.sum
	}

	// Propagate forward if needed
	if s.flipped {
		s.leftSegment.flipRange(s.leftSegment.l, s.leftSegment.r)
		s.rightSegment.flipRange(s.rightSegment.l, s.rightSegment.r)
		s.flipped = !s.flipped
	}

	// Return sum
	return s.leftSegment.sumRange(l, r) + s.rightSegment.sumRange(l, r)
}

func (s *switchSegment) flipRange(l, r int) {
	// Disjoint
	if r < s.l || l > s.r {
		return
	}

	// Covers
	if l <= s.l && s.r <= r {
		s.sum = s.count - s.sum
		s.flipped = !s.flipped
		return
	}

	// Propagate forward if needed
	if s.flipped {
		s.leftSegment.flipRange(s.leftSegment.l, s.leftSegment.r)
		s.rightSegment.flipRange(s.rightSegment.l, s.rightSegment.r)
		s.flipped = false
	}

	// Flip what is needed
	s.leftSegment.flipRange(l, r)
	s.rightSegment.flipRange(l, r)
	s.sum = s.leftSegment.sum + s.rightSegment.sum
}
