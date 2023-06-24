package algorithms

import "src/src/stdlib"

// BinarySearchSimple returns the random index if there are duplicates, and works in O(logn) time, O(1) space
func BinarySearchSimple[T stdlib.Ordered](nums []T, valueToFind T) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] < valueToFind {
			l = m + 1
		} else if nums[m] > valueToFind {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}

// BinarySearchLeftmost returns the leftmost index if there are duplicates, and works in O(logn) time, O(1) space
func BinarySearchLeftmost[T stdlib.Ordered](nums []T, valueToFind T) int {
	l := 0
	r := len(nums) - 1
	foundIndex := -1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] < valueToFind {
			l = m + 1
		} else if nums[m] > valueToFind {
			r = m - 1
		} else {
			foundIndex = m
			r = foundIndex - 1
		}
	}

	return foundIndex
}

// BinarySearchRightmost returns the rightmost index if there are duplicates, and works in O(logn) time, O(1) space
func BinarySearchRightmost[T stdlib.Ordered](nums []T, valueToFind T) int {
	l := 0
	r := len(nums) - 1
	foundIndex := -1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] < valueToFind {
			l = m + 1
		} else if nums[m] > valueToFind {
			r = m - 1
		} else {
			foundIndex = m
			l = foundIndex + 1
		}
	}

	return foundIndex
}
