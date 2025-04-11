package round2

import "fmt"

func wordBreak() {
	input := "leetcode"
	wordDict := []string{"leet", "code"}

	checker := make(map[string]bool)

	for _, word := range wordDict {
		checker[word] = true
	}

	dp := make([]bool, len(input)+1)
	dp[0] = true

	for i := 1; i <= len(input); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && checker[input[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	fmt.Println(dp[len(input)])
}
