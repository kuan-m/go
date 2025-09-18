package main

import "fmt"

func main() {
	wallets := make(map[string]int)

	wallets["Kuan"] = 100000

	delete(wallets, "Kuan")

	if value, exists := wallets["Kuan"]; exists {
		fmt.Println("Баланс :", value)
	} else {
		fmt.Println("Не существует(")
	}
}
