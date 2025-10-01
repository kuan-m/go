package main

import "fmt"

func main() {

	prices := []int{1, 2, 3, 5, 7, 10}

	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {

		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit

}
