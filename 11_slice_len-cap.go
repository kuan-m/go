package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	// drop its first 2 values
	s = s[3:]
	// len – сколько реально видно элементов.
	// cap – сколько максимум можно захватить вперёд до конца исходного массива
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
