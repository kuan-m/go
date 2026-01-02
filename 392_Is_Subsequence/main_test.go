package main

import "testing"

type isSubsequenceTestCase struct {
	name     string
	sub      string
	s        string
	expected bool
}

func TestIsSubsequence(t *testing.T) {
	for _, tt := range getIsSubsequenceTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubsequence(tt.sub, tt.s)

			if result != tt.expected {
				t.Errorf(
					"expected %v, got %v (sub=%q, s=%q)",
					tt.expected,
					result,
					tt.sub,
					tt.s,
				)
			}
		})
	}
}

func getIsSubsequenceTestCases() []isSubsequenceTestCase {
	return []isSubsequenceTestCase{
		{
			name:     "example true",
			sub:      "abc",
			s:        "ahbgdc",
			expected: true,
		},
		{
			name:     "example false",
			sub:      "axc",
			s:        "ahbgdc",
			expected: false,
		},
		{
			name:     "empty subsequence",
			sub:      "",
			s:        "abc",
			expected: true,
		},
		{
			name:     "sub longer than s",
			sub:      "abcd",
			s:        "abc",
			expected: false,
		},
		{
			name:     "same strings",
			sub:      "abc",
			s:        "abc",
			expected: true,
		},
		{
			name:     "repeated characters false",
			sub:      "aaaaaa",
			s:        "bbaaaa",
			expected: false,
		},
		{
			name:     "repeated characters true",
			sub:      "aaa",
			s:        "bbaaaa",
			expected: true,
		},
		{
			name:     "single character true",
			sub:      "a",
			s:        "a",
			expected: true,
		},
		{
			name:     "single character false",
			sub:      "a",
			s:        "b",
			expected: false,
		},
	}
}
