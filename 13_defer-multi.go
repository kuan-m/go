package main

import "fmt"

func main() {

	// Deferred function calls are pushed onto a stack. When a function returns,
	// its deferred calls are executed in last-in-first-out order.
	fmt.Println("Counting: ")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
