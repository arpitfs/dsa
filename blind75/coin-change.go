package blind75

import "fmt"

func coinChange() {
	input := []int{1, 2, 5}
	target := 11
	fmt.Println(minCoins(input, target))
}

func minCoins(input []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 1; i <= target; i++ {
		dp[i] = target + 1
	}

	for _, coin := range input {
		for i := coin; i <= target; i++ {
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[target] == target+1 {
		return -1
	}

	return dp[target]
}
