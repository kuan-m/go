package main

import "fmt"

// A struct is a collection of fields.
type Character struct {
	health int
	attack int
}

func main() {
	char := Character{}
	char_2 := Character{health: 1000, attack: 150}
	char_3 := Character{health: 500}

	fmt.Println(char, char_2, char_3)
	// 			{0 0} {1000 150} {500 0}
}
