package main

import "fmt"

// A struct is a collection of fields.
type Character struct {
	health int
	attack int
}

func main() {
	char := Character{100, 10}
	char_pointer := &char

	char_pointer.health = 1e9
	char_pointer.attack = 1e3

	fmt.Println(char.health) // станет миллиард
	fmt.Println(char.attack) // станет тысяча
}
