package main

import "fmt"

func accessToElement1() {
	data := [3]int{1, 2, 3}

	idx := 4
	fmt.Println(data[idx]) // panic

	fmt.Println(data[4]) // compilation error
}

func accessToElement2() {
	data := [3]int{1, 2, 3}

	idx := -1
	fmt.Println(data[idx]) // panic

	fmt.Println(data[-1]) // compilation error
}

func arrayLen() {
	data := [10]int{}
	fmt.Println(len(data)) // 10 - array never equals to nil
}

func arrayCap() {
	data := [10]int{}
	fmt.Println(len(data)) // same as len
}

func arrayComparasion() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	// except arrays whose element types are incomparable types
	fmt.Println(first == second)
	fmt.Println(first != second)

	// <, <=, >, >= are compilation error
}

func arrayCreation() {
	length1 := 100
	var data1 [length1]int // compilation error
	_ = length1
	_ = data1

	const length2 = 100
	var data2 [length2]int // but const works
	_ = data2
}

func makeArray() {
	_ = make([10]int, 10) // compilation error
}

func appendToArray() {
	_ = append([10]int{}, 10) // compilation error
}

func main() {

	fmt.Println(data1, data2, data3, data4, data5, data6)
}
