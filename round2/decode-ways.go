package round2

import "fmt"

func decodeWays() {
	input := "12"
	fmt.Println(determineWays(input))
}

func determineWays(input string) int {
	if len(input) == 0 || input[0] == '0' {
		return 0
	}

	dp := make([]int, len(input)+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= len(input); i++ {
		oneDigit := input[i-1 : i]
		twoDigit := input[i-2 : i]

		if oneDigit >= "1" && oneDigit <= "9" {
			dp[i] += dp[i-1]
		}
		if twoDigit >= "10" && twoDigit <= "26" {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(input)]
}
