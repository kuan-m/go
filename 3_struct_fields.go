package main

import "fmt"

// A struct is a collection of fields.
type Character struct {
	health int
	attack int
}

func main() {
	char := Character{100, 10}

	fmt.Println(char.health)
	fmt.Println(char.attack)
}
