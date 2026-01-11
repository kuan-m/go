package main

import "fmt"

func accessToElement1() {
	data := make([]int, 3)
	fmt.Println(data[4]) // panic (на этапе компиляции не проверяется)
}

func accessToElement2() {
	data := make([]int, 3, 6)
	fmt.Println(data[4]) // panic
}

func accessToNilSlice1() {
	var data []int
	_ = data[0] // panic
}

func accessToNilSlice2() {
	var data []int
	data[0] = 10 // panic
}

func appendToNilSlice() {
	var data []int
	data = append(data, 10)
}

func rangeByNilSlice() {
	var data []int
	for range data {

	}
}

func makeZeroslice() {
	data := make([]int, 0)
	fmt.Println(len(data)) // 0
	fmt.Println(cap(data)) // 0
}

func makeSlice() {
	_ = make([]int, -5)    // compilation error
	_ = make([]int, 10, 5) // compilation error len can'be longer than cap

	size := -5
	_ = make([]int, size) // panic
}

func realocateSlice() {
	data := []int{1, 2, 3, 4}
	data = append(data, 5)
	// Ёмкость (cap) — это оптимизация:
	// 1. Избегаем частых перевыделений памяти
	// 2. Резервируем место для будущих append операций
	// len - сколько элементов у меня есть
	// cap - "сколько элементов я могу вместить без переезда на новый адрес"

	// What happens:
	// 1. Current capacity (4) is less than needed length (5)
	// 2. Go creates a NEW underlying array with doubled capacity (8)
	// 3. All existing elements [1, 2, 3, 4] are copied to the new array
	// 4. New element 5 is added at index 4
	// 5. The slice 'data' now points to the NEW array: [1, 2, 3, 4, 5]
	// 6. Old array [1, 2, 3, 4] becomes eligible for garbage collection
}
