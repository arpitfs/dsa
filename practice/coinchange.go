package practice

import "fmt"

func coinChange() {
	coins := []int{1, 2, 5}
	amount := 11

	result := minCoin(coins, amount)
	if result == -1 {
		fmt.Println("It is not possible to make the amount with the given coins.")
	} else {
		fmt.Printf("Minimum coins required to make %d is: %d\n", amount, result)
	}
}

func minCoin(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
