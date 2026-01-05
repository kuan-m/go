package main

import "testing"

type backspaceCompareTestCase struct {
	name     string
	s        string
	t        string
	expected bool
}

func getBackspaceCompareTestCases() []backspaceCompareTestCase {
	return []backspaceCompareTestCase{
		{
			name:     "both empty after backspaces",
			s:        "ab##",
			t:        "c#d#",
			expected: true,
		},
		{
			name:     "simple equal",
			s:        "ab#c",
			t:        "ad#c",
			expected: true,
		},
		{
			name:     "simple not equal",
			s:        "a#c",
			t:        "b",
			expected: false,
		},
		{
			name:     "multiple backspaces exceed chars",
			s:        "####",
			t:        "",
			expected: true,
		},
		{
			name:     "backspaces only on one side",
			s:        "a##c",
			t:        "#a#c",
			expected: true,
		},
		{
			name:     "different final length",
			s:        "abc###",
			t:        "a",
			expected: false,
		},
		{
			name:     "no backspaces equal",
			s:        "abc",
			t:        "abc",
			expected: true,
		},
		{
			name:     "no backspaces not equal",
			s:        "abc",
			t:        "abd",
			expected: false,
		},
	}
}

func TestBackspaceCompare(t *testing.T) {
	for _, tt := range getBackspaceCompareTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			result := backspaceCompare(tt.s, tt.t)
			if result != tt.expected {
				t.Fatalf(
					"backspaceCompare(%q, %q) = %v, expected %v",
					tt.s,
					tt.t,
					result,
					tt.expected,
				)
			}
		})
	}
}
