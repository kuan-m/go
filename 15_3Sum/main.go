package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/3sum/description/

// Pattern: Two Pointers
// Time: O(n^2)
// Space: O(1)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {

	// Step 1. Sort
	slices.Sort(nums) // O(nlogn)
	n := len(nums)
	result := [][]int{}

	// Step 2: Fix one element and solve 2Sum for the rest
	for i := 0; i < n; i++ {
		// nums[num] + nums[left] + nums[right] = 0
		// nums[left] + nums[right] = -nums[num]

		// Skip duplicates for the fixed element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Step 3. Two pointers
		target := -nums[i]
		left, right := i+1, n-1

		for left < right {

			sum := nums[left] + nums[right]

			if sum == target {

				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {

				left += 1

			} else {

				right -= 1

			}
		}

	}

	return result
}
