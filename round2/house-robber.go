package round2

import "fmt"

func houseRobber() {
	input := []int{1, 2, 3, 1}
	dp := make([]int, len(input)+1)
	dp[0] = input[0]
	dp[1] = maximum(input[0], input[1])
	for i := 2; i < len(input); i++ {
		rob := input[i] + dp[i-2]
		notRob := dp[i-1]
		dp[i] = maximum(rob, notRob)
	}

	fmt.Println(dp[len(input)-1])
}
