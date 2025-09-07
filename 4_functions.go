package main

import "fmt"

func main() {
	fmt.Print(add(10, 5))
}

func add(x int, y int) int {
	return x + y
}
