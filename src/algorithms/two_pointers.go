package algorithms

/*
TwoPointersSimple is an example of a problem where two pointers are handy.
It works in O(n) time, O(n) space

Problem: https://leetcode.com/problems/squares-of-a-sorted-array/
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in
non-decreasing order.

General pattern:

function fn(arr):

	left = 0
	right = arr.length - 1

	while left < right:
		Do some logic
		Decide how to increase the pointers:
			1. left++
			2. right--
			3. Both left++ and right--
*/
func TwoPointersSimple(arr []int) []int {
	sortedSquares := make([]int, len(arr))

	l := 0
	r := len(arr) - 1

	for c := len(arr) - 1; c >= 0; c-- {
		var biggerSquare int
		if arr[l]*arr[l] > arr[r]*arr[r] {
			biggerSquare = arr[l] * arr[l]
			l++
		} else {
			biggerSquare = arr[r] * arr[r]
			r--
		}

		sortedSquares[c] = biggerSquare
	}

	return sortedSquares
}

/*
TwoPointersDouble is an example of a problem when we iterate on two arrays and where two pointers are handy.
It’s a good template because we don’t deal with the edge cases in the loop

I don't have a problem for it, but it's nice to know!

General pattern:

function fn(arr1, arr2):

	i = 0
	j = 0
	while i < arr1.length AND j < arr2.length:
		Do some logic
		Decide how to increase the pointers:
			1. i++
			2. j++
			3. Both i++ and j++

	// Make sure both iterables are exhausted
	while i < arr1.length:
		Do some logic here depending on the problem
		i++
	while j < arr2.length:
		Do some logic here depending on the problem
		j++
*/
