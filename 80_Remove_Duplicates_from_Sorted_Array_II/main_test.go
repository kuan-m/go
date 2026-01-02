package main

import (
	"reflect"
	"testing"
)

type removeDuplicatesTestCase struct {
	name     string
	input    []int
	expected []int
	k        int
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range getRemoveDuplicatesTestCases() {
		t.Run(tt.name, func(t *testing.T) {

			// делаем копию, чтобы не мутировать оригинал
			nums := append([]int{}, tt.input...)

			k := removeDuplicates(nums)

			if k != tt.k {
				t.Errorf("expected k=%d, got %d", tt.k, k)
			}

			if !reflect.DeepEqual(nums[:k], tt.expected) {
				t.Errorf(
					"expected nums[:k]=%v, got %v",
					tt.expected,
					nums[:k],
				)
			}
		})
	}
}

func getRemoveDuplicatesTestCases() []removeDuplicatesTestCase {
	return []removeDuplicatesTestCase{
		{
			name:     "example from leetcode",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: []int{1, 1, 2, 2, 3},
			k:        5,
		},
		{
			name:     "all elements same",
			input:    []int{2, 2, 2, 2},
			expected: []int{2, 2},
			k:        2,
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
			k:        4,
		},
		{
			name:     "only two duplicates",
			input:    []int{1, 1},
			expected: []int{1, 1},
			k:        2,
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
			k:        1,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
			k:        0,
		},
		{
			name:     "mixed negatives",
			input:    []int{-1, -1, -1, 0, 0, 1},
			expected: []int{-1, -1, 0, 0, 1},
			k:        5,
		},
	}
}
