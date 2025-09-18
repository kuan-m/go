package main

import "fmt"

func main() {
	names := [4]string{
		"Leonardo",
		"Donatello",
		"Raphael",
		"Mikey",
	}
	fmt.Println(names)

	a := names[0:2]   // [Leonardo Donatello]
	b := names[2:]    // or names[2:4] [Leonardo Donatello]
	fmt.Println(a, b) //  [Leonardo Donatello] [Raphael Mikey]

	a[0] = "XXX" // changing the elements of a slice modifies the elements of its underlying array
	b[0] = "XXX"
	fmt.Println(a, b) // [XXX Donatello] [XXX Mikey]
	// a slices does not store any data, it just descirbes a section of an underlying array
}
