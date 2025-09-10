package main

import "fmt"

func main() {

	// defer отложенный вызов, вызывается после окончания кода.
	// по очереди LIFO
	defer fmt.Println("Третий")
	defer fmt.Println("Второй")
	defer fmt.Println("Первый")

	fmt.Println("Конец")
}
