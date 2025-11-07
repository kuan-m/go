package main

import "fmt"

func factorialIterative(n int) int {
	ans := 1
	for n > 0 {
		ans *= n
		n -= 1
	}

	return ans
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func main() {
	fmt.Println(factorialIterative(5))
	fmt.Println(factorialRecursive(5))
	// │
	// ├─ 5 * factorialRecursive(4)
	// │  │
	// │  ├─ 4 * factorialRecursive(3)
	// │  │  │
	// │  │  ├─ 3 * factorialRecursive(2)
	// │  │  │  │
	// │  │  │  ├─ 2 * factorialRecursive(1)
	// │  │  │  │  │
	// │  │  │  │  ├─ 1 * factorialRecursive(0)
	// │  │  │  │  │  │
	// │  │  │  │  │  └─ return 1  ← Базовый случай!
	// │  │  │  │  │
	// │  │  │  │  └─ return 1
	// │  │  │  │
	// │  │  │  └─ return 2
	// │  │  │
	// │  │  └─ return 6
	// │  │
	// │  └─ return 24
	// │
	// └─ return 120
}
