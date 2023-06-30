# Sorting

Sorting is good for subsequences or subsets, because this allows you to binary search.
Also it then improves iteration, you know that each next value will be higher, so you can leverage that somehow.

```go
nums := []int{7, 4, 2, 5, 9, 1}

// Sort ascending
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

fmt.Println(nums) // [1 2 4 5 7 9]
```
