package main

import "fmt"

func main() {
	fmt.Println(split(41412))
}

func split(sum int) (x, y int) {
	x = sum / 2
	y = x / 2
	return
}
