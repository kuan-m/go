package main

import (
	"reflect"
	"testing"
)

type bubbleSortTestCase struct {
	name     string
	nums     []int
	expected []int
}

func TestBubbleSort(t *testing.T) {
	tests := []bubbleSortTestCase{
		{
			name:     "simple from 5 to 1",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "already sorted",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "with duplicates",
			nums:     []int{3, 1, 2, 3, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "two elements reversed",
			nums:     []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "single element",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "empty slice",
			nums:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := bubbleSort(tt.nums)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf(
					"bubbleSort(%v) = %v; expected %v",
					tt.nums,
					result,
					tt.expected,
				)
			}
		})
	}
}
