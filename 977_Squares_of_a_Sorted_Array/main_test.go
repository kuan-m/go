package main

import (
	"reflect"
	"testing"
)

func TestSquaresOfSortedArray(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}

	expected := []int{0, 1, 9, 16, 100}

	result := sortedSquares(nums)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
