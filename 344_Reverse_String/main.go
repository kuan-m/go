package main

import "fmt"

// https://leetcode.com/problems/reverse-string/description/

// Pattern: Two Pointers
// Time: O(n) Space: O(1)

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(s)

	fmt.Print(s)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
