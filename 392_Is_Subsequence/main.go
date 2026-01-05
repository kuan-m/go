package main

import (
	"fmt"
)

// https://leetcode.com/problems/is-subsequence/description/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(1)

func main() {
	sub := "abc"
	s := "ahbgdc"

	fmt.Println(isSubsequence(sub, s))
}

func isSubsequence(sub string, s string) bool {
	p1, p2 := 0, 0

	for p1 < len(sub) && p2 < len(s) {

		if sub[p1] == s[p2] {
			p1++
		}

		p2++
	}

	return p1 == len(sub)
}
