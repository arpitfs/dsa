package round2

import "fmt"

func stock() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices)-1; i++ {
		currentProfit := prices[i] - minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	fmt.Println(maxProfit)
}
