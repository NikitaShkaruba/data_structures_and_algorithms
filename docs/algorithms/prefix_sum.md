# Prefix sum

Prefix sum precomputes sums of arrays, letting us later do quick range queries for `O(1)` time instead of usual `O(n)` time.
The one downside of this pattern is extra `O(n)` space, but it's well worth it, extra space is okay for most solutions.
The next level of this pattern is [Segment tree](../../src/data_structures/segment_tree.go).

```go
// nums       = [4, 2, 3,  5,  2]
// prefixSums = [4, 6, 9, 14, 16]
//                  l      r   
// getSubarraySum(prefixSums, 1, 3) == 10

// Works in O(n) time, O(n) space
func createPrefixSum(nums []int) []int {
	prefixSums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefixSums = append(prefixSums, prefixSums[i-1] + nums[i])
	}
	return prefixSums
}

// Basic version, works in O(1) time, O(1) space
func getSubarraySum(prefixSums []int, l, r int) int {
	if l == 0 {
		return prefixSums[r]
	}
	
	return prefixSums[r] - prefixSums[l-1]
}

// More convenient one-liner, but requires nums, works in O(1) time, O(1) space
func getSubarraySum(nums, prefixSums []int, l, r int) int {
	return prefixSums[r] - prefixSums[l] + nums[l]
}
```

How to get range sum without prefix sum
```go
// Basic version, works in O(n) time, O(1) space
func getSubarraySum(nums []int, l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += nums[i]
	}
	return sum
}

```
