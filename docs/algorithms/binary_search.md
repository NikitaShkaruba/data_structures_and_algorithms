# Binary search

Binary search helps to find value indexes in sorted array in `O(logn)` time, `O(1)` space.  
It can be implemented very differently, but I have this template. Its flawless, you should use it everywhere.

### With binary search

```go
// Works in O(logn) time, O(1) space
func binarySearch(arr []int, valueToFind int) int {
	l := 0
	r := len(arr) - 1
	foundIndex := -1

	for l <= r {
		m := l + (r-l)/2 // Less overflow-prone way of finding the middle

		// Change the conditions depending on the problem
		if arr[m] < valueToFind {
			l = m + 1
		} else if arr[m] > valueToFind {
			r = m - 1
		} else {
			foundIndex = m
			// Continue searching for the better value. You would want to change this line
			r = foundIndex - 1
		}
	}

	return foundIndex
}

```

### Without binary search

Without it, we need to iterate over all the data in `O(n)` time, which is a huge overhead compared to `O(logn)` time

```go
// Works in O(n) time, O(1) space
func withoutBinarySearch(arr []int, valueToFind int) int {
	for i, v := range arr {
		if v == valueToFind {
			return i
		}
	}
	return -1
}
```
