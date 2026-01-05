package main

import "fmt"

func main() {

	fmt.Println(gfc(12, 20))
}

func gfc(num1, num2 int) int {
	// 20 = 1,2,4,5,10,20
	// 12 = 1,2,3,4,6,12

	smaller := min(num1, num2)
	ans := 1

	for i := 1; i <= smaller+1; i++ {
		if num1%i == 0 && num2%i == 0 {
			ans = i
		}
	}

	return ans
}
