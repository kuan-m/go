package main

import "fmt"

// A struct is a collection of fields.
type Character struct {
	health int
	attack int
}

func main() {
	fmt.Println(Character{100, 10})
}
