package dp

import (
	"fmt"
	"math"
)

func coins() {
	coins := []int{3, 5, 6}
	amount := 6
	numberOfCoins := getCoins(coins, amount)
	fmt.Println(numberOfCoins)
}

func getCoins(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		var min int = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != -1 {
				min = minimum(min, dp[i-coins[j]])
			}
		}
		if min == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}

	return dp[amount]
}

func minimum(value1, value2 int) int {
	if value1 < value2 {
		return value1
	}
	return value2
}
