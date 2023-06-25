# Sliding window

Helps to not recreate an array of subset data

### Sliding window of undefined size

```go
func slidingWindow(arr []int) {
	l := 0

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
	// Build the initial window
	for i := 0; i < size; i++ {
		// Logic that "adds" arr[i] to the window
	}

	// You have a correct window state, you might as well use it depending on the problem

	// Move the window
	for i := size; i < len(arr); i++ {
		// Logic that "adds" arr[i] to the window
		// Logic that "removes" arr[i-size] from the window
		// You have a correct window state, you might as well use it depending on the problem
	}
}

```
