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

### Space / answer solutions

A very common type of problem is "what is the max/min that something can be done". Binary search can be used if the following criteria are met:

- You can quickly (in O(n) or better) verify if the task is possible for a given number x.
- If the task is possible for a number x, and you are looking for:
  - A maximum, then it is also possible for all numbers less than x.
  - A minimum, then it is also possible for all numbers greater than x.
- If the task is not possible for a number x, and you are looking for:
  - A maximum, then it is also impossible for all numbers greater than x.
  - A minimum, then it is also impossible for all numbers less than x.
  
The 2nd and 3rd requirements imply that there are two "zones". One where it is possible and one where it is impossible. The zones have no breaks, no overlap, and are separated by a threshold.

When a problem wants you to find the min/max, it wants you to find the threshold where the task transitions from impossible to possible. First, we establish the possible solution space by identifying the minimum possible answer and the maximum possible answer. Next, we binary search on this solution space. For each mid, we perform a check to see if the task is possible. Depending on the result, we halve the search space. Eventually, we will find the threshold.

Solutions like this is only viable if your time complexity fits, so you need to compute if this time complexity works for you.
