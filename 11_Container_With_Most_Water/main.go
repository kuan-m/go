package main

import (
	"fmt"
)

// https://leetcode.com/problems/container-with-most-water/description/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(1)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// 49 = width(7) * minHeight(7)
	fmt.Println(maxArea(nums))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	area := 0
	maxArea := 0
	minHeight := 0

	for left < right {

		width := right - left

		minHeight = min(height[left], height[right])

		area = width * minHeight

		maxArea = max(maxArea, area)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
