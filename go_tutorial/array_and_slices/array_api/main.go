package main

import "fmt"

func main() {
	data := [5]int{2: 5, 6, 1: 7} // [0 7 5 6 0]

	fmt.Println(data[4])   // read from array
	data[4] = 10000        // write into array
	fmt.Println(len(data)) // get array length = 5
	fmt.Println(cap(data)) // get array cap = 5
	p := &data             // points to array
	fmt.Printf("%p\n", p)  //  0xc0000a6000
	fmt.Println(*p)        // [0 7 5 6 10000]
	s := data[1:4]         // get slice from array
	fmt.Println(s)
}
