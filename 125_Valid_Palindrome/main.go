package main

import (
	"fmt"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/

// Pattern: Two Pointers
// Time: O(n) Space: O(1)

func main() {
	str := "aboba"
	fmt.Print(isPalindrome(str))
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		l := rune(s[left])
		r := rune(s[right])

		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
		} else if unicode.ToLower(l) == unicode.ToLower(r) {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
