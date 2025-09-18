package main

import "fmt"

type Character struct {
	health, attack int
}

var chars = map[string]Character{
	"Kuan": Character{
		10000, 100,
	},
	"Boss": Character{
		100000, 1000,
	},
}

func main() {

	fmt.Println(chars)

}
