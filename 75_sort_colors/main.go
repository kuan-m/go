package main

import "fmt"

func main() {
	numbers := []int{2, 0, 2, 1, 1, 0}
	sortColors(numbers)
	fmt.Println(numbers)
}

func sortColors(nums []int) {
	myMap := map[int]int{}

	for _, v := range nums {
		myMap[v]++
	}
	fmt.Println(myMap) // [0:2 1:2 2:2

	idx := 0
	for i := 0; i <= 2; i++ {
		for j := 0; j < myMap[i]; j++ {
			nums[idx] = i
			idx++
		}
	}
}
