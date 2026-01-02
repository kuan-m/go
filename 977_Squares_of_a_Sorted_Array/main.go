package main

import (
	"fmt"
)

// https://leetcode.com/problems/squares-of-a-sorted-array/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(n)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	// 0, 1, 9, 16, 100
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	result := make([]int, right+1)
	pos := right

	for left <= right {
		multipliedLeft := nums[left] * nums[left]
		multipliedRight := nums[right] * nums[right]

		if multipliedLeft > multipliedRight {
			result[pos] = multipliedLeft
			left++
		} else {
			result[pos] = multipliedRight
			right--
		}
		pos--
	}

	return result
}
