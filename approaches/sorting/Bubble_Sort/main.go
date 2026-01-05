package main

import "fmt"

// Time: O(n^2)
// Space: O(1)

func main() {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(bubbleSort(nums))

}

// Удобство? Прост в реализации.
func bubbleSort(nums []int) []int {

	n := len(nums) - 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
