package main

import "fmt"

func main() {
	// so arrays cannot be resized.
	var arr [3]string
	arr[0] = "Sanya"
	arr[1] = "Petya"
	arr[2] = "Katya"
	fmt.Println("Names: ", arr)
	fmt.Println(arr[0], arr[1])

	digits := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(digits)
}
