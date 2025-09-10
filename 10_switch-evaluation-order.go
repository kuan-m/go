package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Saturday?")

	today := time.Now().Weekday()
	switch time.Thursday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow.")
	case today + 2:
		fmt.Println("It two days.")
	default:
		fmt.Println("Too far away.")
	}
}
