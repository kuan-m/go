package main

import "fmt"

func modifyValue(data int) {
	data = 5
}

func modifyPointer(data *int) {
	*data = 1
}

func main() {

	// Без указателей функции работают с копиями:
	x := 100
	modifyValue(x)    // останется 100
	modifyPointer(&x) // x станет 1
	fmt.Println(x)

	i := 101
	p := &i

	fmt.Println(i)
	fmt.Println(p)  // points to I (saves the adress 0xc00000a0e8)
	fmt.Println(*p) // read i thrpugh the pointer

	*p = 102       // set i through the pointer p
	fmt.Println(i) // see the new value of j

}
