package main

import "fmt"

func main() {
	// Variables declared without an explicit initial value are given their zero value.
	var i int     // 0
	var f float64 // 0
	var b bool    // false
	var s string  // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
