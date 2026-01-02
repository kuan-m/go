package main

import (
	"testing"
)

func TestMaxArea_EdgeCases(t *testing.T) {
	for _, tt := range getMaxAreaEdgeCases() {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.height)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func getMaxAreaEdgeCases() []struct {
	name     string
	height   []int
	expected int
} {
	return []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "minimum input",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "all zeros",
			height:   []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "all same heights",
			height:   []int{5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "strictly increasing",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "strictly decreasing",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6,
		},
		{
			name:     "max area in the middle",
			height:   []int{1, 100, 1, 1, 100, 1},
			expected: 300,
		},
		{
			name:     "max heights at edges",
			height:   []int{10000, 1, 1, 1, 10000},
			expected: 40000,
		},
	}
}
