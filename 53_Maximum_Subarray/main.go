package main

import "fmt"

func main() {
	// https://leetcode.com/problems/maximum-subarray/description/
	// Time: O(n) Memory: O(1)
	data := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(data))
}

func maxSubArray(nums []int) int {
	ans, temp := nums[0], 0

	for _, val := range nums {
		temp += val
		ans = max(temp, ans)
		temp = max(0, temp)
	}

	return ans
}
