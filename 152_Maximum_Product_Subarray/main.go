package main

import "fmt"

func main() {
	// https://leetcode.com/problems/maximum-product-subarray/
	// Time: O(n) Memory: O(1)
	nums := []int{2, 3, -2, 4}

	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxP, minP, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxP, minP = minP, maxP
		}

		maxP = max(nums[i], maxP*nums[i])
		minP = min(nums[i], minP*nums[i])

		ans = max(ans, maxP)
	}

	return ans
}
