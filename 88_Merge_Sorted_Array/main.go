package main

import "fmt"

// https://leetcode.com/problems/merge-sorted-array/description/

// Time: O(n+m)
// Memory: O(1)

func main() {
	nums1 := []int{1, 3, 4, 0, 0, 0}
	nums2 := []int{3, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1

	k := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[k] = nums1[p1]
			p1--
		} else {
			nums1[k] = nums2[p2]
			p2--
		}
		k--
	}

	// если в nums2 остались элементы
	for p2 >= 0 {
		nums1[k] = nums2[p2]
		p2--
		k--
	}
}
