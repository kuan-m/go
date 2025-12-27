package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}

	moveZeroes(nums)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected %v, got %v", expected, nums)
	}
}
