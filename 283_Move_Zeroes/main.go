package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/description/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(1)

func main() {

	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

	fmt.Println(nums)
}

func moveZeroes(nums []int) {

	slowP := 0

	for fastP, num := range nums {
		if num != 0 {
			nums[slowP] = num
			if slowP != fastP {
				nums[fastP] = 0
			}
			slowP++
		}
	}

}
