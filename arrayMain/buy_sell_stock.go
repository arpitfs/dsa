package arrayMain

import "fmt"

func buy_sell() {
	prices := []int{7, 1, 5, 3, 6, 4}
	miniPrice := prices[0]
	maxprofit := 0
	for i := 1; i < len(prices)-1; i++ {
		profit := prices[i] - miniPrice

		if prices[i] < miniPrice {
			miniPrice = prices[i]
		}
		if profit > maxprofit {
			maxprofit = profit
		}
	}

	fmt.Println("Max Profit", maxprofit)

}
