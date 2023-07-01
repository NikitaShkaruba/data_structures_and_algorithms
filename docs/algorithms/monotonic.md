Monotonic data structure is an array-like structure that is always sorted - this property helps to keep good time complexity by not iterating over all the elements.

Monotonic stacks and queues are useful in problems that, for each element, involves finding the "next" element based on some criteria, for example, the next greater element.
They're also good when you have a dynamic window of elements, and you want to maintain knowledge of the maximum or minimum element as the window changes.
In more advanced problems, sometimes a monotonic stack or queue is only one part of the algorithm.

In the snippet below we keep all the temperatures monotonically sorted, which helps us to not iterate over all the values in the stack.
We don't need to iterate over all the data, because it's sorted, and if the last stack temperature is not lower than the current temperature, next stack temperatures will only be greater, so we can prematurely stop.
```go
// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have to wait
// after the ith day to get a warmer temperature.
// If there is no future day for which this is possible, keep answer[i] == 0 instead.
//
// https://leetcode.com/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	warmerTemperatures := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := range temperatures {
		// Handle temperatures that are lower then ours
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			warmerTemperatures[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return warmerTemperatures
}
```
