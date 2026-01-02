package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(1)

func main() {

	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	// 1. подход с памятью O(n)
	// var seen = map[int]int{}

	// k := 0
	// for _, num := range nums {
	// 	if seen[num] < 2 {
	// 		nums[k] = num
	// 		seen[num]++
	// 		k++
	// 	}
	// }

	// 2. подход с памятью O(1)
	k := 0

	for _, num := range nums {
		if k < 2 || num != nums[k-2] {
			nums[k] = num
			k++
		}
	}

	fmt.Println(nums[:k])
	return len(nums[:k])
}
