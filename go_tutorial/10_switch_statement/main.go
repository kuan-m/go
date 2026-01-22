package main

import (
	"fmt"
)

func main() {

	finger := 4
	fmt.Printf("Finger %d is ", finger)

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Incorrect finger number")
	}

	// multiple expressions
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Printf("%s is a vowel", letter)
	default:
		fmt.Printf("%s is not a vowel", letter)
	}

}
