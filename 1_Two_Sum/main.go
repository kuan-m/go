package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, 9))

}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for idx, val := range nums {
		x := target - val

		if y, exists := seen[x]; exists {
			return []int{y, idx}
		}

		seen[val] = idx
	}

	return nil
}
