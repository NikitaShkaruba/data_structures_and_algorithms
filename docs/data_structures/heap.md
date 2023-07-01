# Heap

If you need to find min or max very fast - try heap.
Heap can be built with `O(n)` time, and you can push \ pop data for `O(logn)` time and peek the max for `O(1)` time.

Source code: [src/data_structures/heap.go](../../src/data_structures/heap.go)

### Common heap usage

In general you put data to heap, and then you can retrieve minimums or maximums very quickly

```go
func heapSort(nums []int) []int {
	h := data_structures.NewEmptyHeap(func(a, b int) bool {
		return a < b // min heap
	})

	h.Push(143)
	h.Push(1)
	h.Push(234)

	res := make([]int, 0)
	for h.GetSize() != 0 {
		res = append(res, h.Pop())
	}
	return res // [1, 234, 143]
}
```

### Reversed heap usage

Heap can be reversed for keeping `k` max / min elements.
It's pretty good, because you don't store excess min / max elements, you only store `k` of them.

```go
func getKMin() []int {
	nums := []int{43, 4, 234, 1, 532}
	k := 2

	h := data_structures.NewEmptyHeap(func(a, b int) bool {
		return a > b // max heap (even though we need k mins)
	})

	for i := range nums {
		h.Push(nums[i])
		if h.GetSize() > k {
			// Remove biggest
			h.Pop()
		}
	}

	res := make([]int, k)
	for i := range res {
		res[i] = h.Pop()
	}

	return res // [4, 1]
}
```

### AVL tree comparison

Heap is very similar to [avl tree](./avl_tree.md), but allows you to find min max for `O(1)` time instead of `O(logn)`,
and can only work with either min or max data, when avl tree can keep track of both min and max.
Also heap can find an element inside of itself for `O(n)` time, when avl tree finds it in `O(logn)` time.

In general avl tree has more usages and better overall, but heap is very simple. So if you consider one of those data structures,
don't forget to consider the other.

