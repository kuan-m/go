package main

import "fmt"

// https://leetcode.com/problems/remove-element/description/

// Pattern: Two Pointers
// Time: O(n) Space: O(1)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(nums, 2)

	fmt.Print(nums)
}

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
