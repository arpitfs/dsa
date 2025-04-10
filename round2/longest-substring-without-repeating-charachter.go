package round2

import "fmt"

func longestSubstringWRC() {
	input := "abcabcbb"
	ans := 0
	left := 0
	checker := make(map[rune]int)

	for i, v := range input {
		if _, e := checker[v]; !e {
			checker[v] = i
		} else {
			left = maximum(left, checker[v]+1)
			checker[v] = i
		}
		ans = maximum(ans, i-left+1)
	}

	fmt.Println(ans)
}
