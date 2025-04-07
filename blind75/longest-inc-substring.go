package blind75

import "fmt"

func findLongestIncreasingSubstring() {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(input))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maximum(dp[i], dp[j]+1)
			}
		}
	}

	r := 0
	for _, v := range dp {
		r = maximum(r, v)
	}

	return r
}
