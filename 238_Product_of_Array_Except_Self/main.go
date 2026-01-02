package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/description/

func productExceptSelf(nums []int) []int {
	len := len(nums)

	ans := []int{}
	for i := 0; i < len; i++ {
		ans = append(ans, 1)
	}

	prefix := 1
	for i := 0; i < len; i++ {
		ans[i] = prefix
		prefix = prefix * nums[i]
	}
	// [1 1 2 6]

	suffix := 1
	for i := len - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix = suffix * nums[i]
	}
	// [24 12 4 1]

	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
