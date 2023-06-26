# Sliding window

Helps to not recreate an array of subset data, but to reuse the currently running array.
Without it moving to the next window is `O(n)` time

### Sliding window of undefined size

```go
func slidingWindow(arr []int) {
	l := 0

	// Each window move is O(1) time, O(1) space
	for r := 0; r < len(arr); r++ {
		// Logic that "adds" arr[r] to the window

		for l < r && isWindowIncorrect() {
			// Logic that "removes" arr[l] from the window
			l++
		}

		// You have a correct window state, you might as well use it depending on the problem
	}
}
```

### Sliding window of fixed size

```go
func slidingWindow(arr []int, size int) {
	// Build the initial window, takes O(size) time, O(1) space
	for i := 0; i < size; i++ {
		// Logic that "adds" arr[i] to the window
	}

	// You have a correct window state, you might as well use it depending on the problem

	// Move the window. Each window move is O(1) time, O(1) space
	for i := size; i < len(arr); i++ {
		// Logic that "adds" arr[i] to the window
		// Logic that "removes" arr[i-size] from the window
		// You have a correct window state, you might as well use it depending on the problem
	}
}
```

### Without sliding window

If the size is huge, we get a big overhead on each window move which can add up to `O(n^2)` time

```go
func withoutSlidingWindow(arr []int, size int) {
	for i := 0; i < len(arr); i++ {
		// Build the window in O(size) time, O(1) space
		for j := i; j < size; j++ {
			// Logic that "adds" arr[i] to the window
		}

		// You have a correct window state, you might as well use it depending on the problem
	}
}
```
