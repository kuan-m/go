package main

import "fmt"

// https://leetcode.com/problems/backspace-string-compare/description/

// Pattern: Two Pointers
// Time: O(n)
// Space: O(1)

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println(backspaceCompare(s, t))
}

func backspaceCompare(s, t string) bool {

	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {

		// двигаем i до валидного символа
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}

		// двигаем j до валидного символа
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		// сравнение
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			// один закончился, другой нет
			return false
		}

		i--
		j--
	}

	return true
}
