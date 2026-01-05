package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "basic merge",
			nums1:  []int{1, 3, 4, 0, 0, 0},
			m:      3,
			nums2:  []int{3, 5, 6},
			n:      3,
			expect: []int{1, 3, 3, 4, 5, 6},
		},
		{
			name:   "nums2 empty",
			nums1:  []int{1, 2, 3},
			m:      3,
			nums2:  []int{},
			n:      0,
			expect: []int{1, 2, 3},
		},
		{
			name:   "nums1 empty",
			nums1:  []int{0, 0, 0},
			m:      0,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{2, 5, 6},
		},
		{
			name:   "all nums2 smaller",
			nums1:  []int{4, 5, 6, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 2, 3},
			n:      3,
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)

			if !reflect.DeepEqual(tt.nums1, tt.expect) {
				t.Errorf(
					"merge() = %v, want %v",
					tt.nums1,
					tt.expect,
				)
			}
		})
	}
}
