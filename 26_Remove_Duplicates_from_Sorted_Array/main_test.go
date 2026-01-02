package main

import (
	"reflect"
	"testing"
)

type removeDuplicatesITestCase struct {
	name     string
	input    []int
	expected []int
	k        int
}

func TestRemoveDuplicatesI(t *testing.T) {
	for _, tt := range getRemoveDuplicatesITestCases() {
		t.Run(tt.name, func(t *testing.T) {

			// копируем вход, чтобы тесты были изолированы
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

func getRemoveDuplicatesITestCases() []removeDuplicatesITestCase {
	return []removeDuplicatesITestCase{
		{
			name:     "example from leetcode",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
			k:        5,
		},
		{
			name:     "already unique",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
			k:        4,
		},
		{
			name:     "all same",
			input:    []int{2, 2, 2, 2},
			expected: []int{2},
			k:        1,
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
			name:     "negative numbers",
			input:    []int{-3, -3, -2, -1, -1, 0},
			expected: []int{-3, -2, -1, 0},
			k:        4,
		},
	}
}
