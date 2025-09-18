package main

import "fmt"

func main() {
	// an array has fixed size
	population := [6]int{1000, 200000, 110000, 434500, 64645, 135667}

	// lsices dynamicaly sized
	var s []int = population[1:4]
	fmt.Println(s)
}
