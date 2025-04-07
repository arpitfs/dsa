package blind75

import "fmt"

func decodeWays() {
	input := "226"
	fmt.Println(decodeWaysLogic(input))
}

func decodeWaysLogic(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < n; i++ {
		oneDigit := s[i-1 : i]
		twoDigit := s[i-2 : i]

		if oneDigit >= "0" && oneDigit <= "9" {
			dp[i] += dp[i-1]
		}

		if twoDigit >= "10" && twoDigit <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
