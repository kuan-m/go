package main

import "fmt"

func main() {
	// 	The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.

	var s = []int{543, 565, 85634, 7543, 6654, 99312}

	for i, v := range s {
		fmt.Printf("i: %d v: %d \n", i, v)
	}

	for i := range s {
		fmt.Printf("index: %d \n", i)
	}

	for _, v := range s {
		fmt.Printf("value: %d \n", v)
	}
}
