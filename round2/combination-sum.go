package round2

import "fmt"

func combinationSum() {
	input := []int{1, 2, 3}
	target := 4

	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range input {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	fmt.Println(dp[target])
}
