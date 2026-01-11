package main

import "fmt"

func main() {
	data := []int{}            // []               	len=0 cap=0
	data1 := []int{1, 2}       // [1 2]            	len=2 cap=2
	data2 := []int{5: 10}      // [0 0 0 0 0 10]   	len=6 cap=6
	data3 := make([]int, 6)    // [0 0 0 0 0 0] 	len=6 cap=6
	data4 := make([]int, 3, 6) // [0 0 00] 			len=3 cap=6

	fmt.Println(data, data1, data2, data3, data4)
}
