package blind75

import "fmt"

func combinationSum() {
	nums := []int{1, 2, 3}
	target := 4

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	fmt.Println(dp[target])
}
