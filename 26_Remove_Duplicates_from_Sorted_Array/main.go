package main

import (
	"fmt"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(1)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
