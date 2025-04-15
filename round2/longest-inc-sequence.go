package round2

import "fmt"

func longestIncreasingSequence() {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//input := []int{1, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6}

	dp := make([]int, len(input))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] < input[i] {
				dp[i] = maximum(dp[i], dp[j]+1)
			}
		}
	}

	r := 0
	for _, v := range dp {
		if v > r {
			r = maximum(r, v)
		}
	}

	fmt.Println(r)

}
