package main

import "fmt"

// https://leetcode.com/problems/contains-duplicate/description/

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int, 0)

	for _, value := range nums {
		if _, isExists := hashMap[value]; isExists {
			hashMap[value]++
			if hashMap[value] >= 2 {
				return true
			}
		} else {
			hashMap[value] = 1
		}
	}

	return false
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
