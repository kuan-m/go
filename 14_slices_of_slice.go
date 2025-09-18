package main

import "fmt"

func main() {
	// Slices can contain any type, including other slices.

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
	}
}
