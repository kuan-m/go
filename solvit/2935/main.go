package main

import (
	"log"
)

// https://solvit.space/coding/2935
// Даны два целочисленных массива nums1 и nums2. Верните массив их пересечения.
// Каждый элемент в результате должен появляться столько раз, сколько он встречается в обоих массивах.
// Можно вернуть результат в любом порядке.

// Пример 1:
// Вход:
// nums1 = [1,2,2,1], nums2 = [2,2]
// Выход:
// [2,2]

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	log.Println(intersect(nums1, nums2))

}

func intersect(nums1 []int, nums2 []int) []int {

	seenNums := make(map[int]int)
	ans := []int{}

	for _, value := range nums2 {
		_, exists := seenNums[value]
		if exists {
			seenNums[value]++
		} else {
			seenNums[value] = 1
		}
	}

	for _, value := range nums1 {
		_, exists := seenNums[value]
		if exists {
			if seenNums[value] > 0 {
				seenNums[value]--
				ans = append(ans, value)
			}
		}
	}

	return ans
}
