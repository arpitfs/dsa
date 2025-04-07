package blind75

import "fmt"

func wordBreak() {
	dict := []string{"leet", "code"}
	s := "leetcode"
	fmt.Println(wordBreakLogic(s, dict))
}

func wordBreakLogic(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)

	for _, v := range wordDict {
		wordSet[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
