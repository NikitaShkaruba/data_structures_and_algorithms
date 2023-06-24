# Two pointers

Two pointers is used for array iteration.  
Usually works in `O(n)` time, `O(1)` space.

### On a single array

```go
func twoPointers(arr []int) {
	l := 0
	r := len(arr) - 1

	for l < r {
		// Do some logic

		// Provide conditions here to decide
		// which pointers should be incremented:
		if isFirstPointerOnly() {
			l++
		} else if isSecondPointerOnly() {
			r--
		} else {
			l++
			r--
		}
	}
}

```

### On two arrays

This template is good because we don’t deal with the edge cases in the loop

```go
func twoPointers(arr1, arr2 []int) {
	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		// Do some logic

		// Provide conditions here to decide
		// which pointers should be incremented:
		if isFirstPointerOnly() {
			i++
		} else if isSecondPointerOnly() {
			j++
		} else {
			i++
			j++
		}
	}

	// Make sure both pointers are exhausted
	for i < len(arr1) {
		// Do some logic
		i++
	}
	for j < len(arr2) {
		// Do some logic
		j++
	}
}

```
