package main

import "fmt"

type Character struct {
	health, attack int
}

func main() {

	characters := make(map[string]Character)
	wallets := make(map[string]int)

	fmt.Println(characters) // map[]

	characters["Kuan"] = Character{
		100, 10,
	}
	wallets["kuan"] = 100000

	fmt.Println(characters) // map[Kuan:{100 10}]
	fmt.Println(wallets)

}
