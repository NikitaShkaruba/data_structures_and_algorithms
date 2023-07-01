# Indexed heap

Basically the same as [heap](./heap.md), but adds a new method of removing by value for `O(logn)` time instead of `O(n)` time,
and a method to search for a value for `O(1)` time instead of `O(n)` time.

Source code: [src/data_structures/indexed_heap.go](../../src/data_structures/indexed_heap.go)

### Finding sliding window median with two indexed heaps

Two indexed heaps can help you find sliding window medians by giving data to each other.

```go
// https://leetcode.com/problems/sliding-window-median/
func medianSlidingWindow(nums []int, k int) []float64 {
	leftHeap := data_structures.NewEmptyIndexedHeap(func(a, b int) bool {
		return a > b // max heap
	})
	rightHeap := data_structures.NewEmptyIndexedHeap(func(a, b int) bool {
		return a < b // min heap
	})

	medians := make([]float64, 0, len(nums)-k+1)

	// Create the first window
	for i := 0; i < k; i++ {
		leftHeap.Push(nums[i])
		rightHeap.Push(leftHeap.Pop())
		// Always have more values in the left heap
		if leftHeap.GetSize() < rightHeap.GetSize() {
			leftHeap.Push(rightHeap.Pop())
		}
	}
	medians = append(medians, getMedian(leftHeap, rightHeap))

	for i := k; i < len(nums); i++ {
		// Remove the minimum
		removedFromLeftHeap := leftHeap.RemoveByValue(nums[i-k])
		if !removedFromLeftHeap {
			rightHeap.RemoveByValue(nums[i-k])
		}

		leftHeap.Push(nums[i])
		rightHeap.Push(leftHeap.Pop())
		// Always have more values in the left heap
		if leftHeap.GetSize() < rightHeap.GetSize() {
			leftHeap.Push(rightHeap.Pop())
		}

		medians = append(medians, getMedian(leftHeap, rightHeap))
	}

	return medians
}

func getMedian(leftHeap *data_structures.IndexedHeap[int], rightHeap *data_structures.IndexedHeap[int]) float64 {
	// If we have even amount of numbers, we should divide two closes to the center together
	if leftHeap.GetSize() == rightHeap.GetSize() {
		return (float64(leftHeap.Peek()) + float64(rightHeap.Peek())) / float64(2)
	}

	return float64(leftHeap.Peek())
}
```
