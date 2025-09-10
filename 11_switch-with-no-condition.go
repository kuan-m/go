package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	// Switch without a condition is the same as switch true.
	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning!")
	case time.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening.")
	}
}
