package string

import "fmt"

func longestSubstringWithoutReptation() {
	input := "abba"
	fmt.Println(longestSubstring(input))
	input = "pwwkew"
	fmt.Println(longestSubstring(input))
}

func longestSubstring(input string) int {
	ans := 0
	left := 0
	checker := make(map[rune]int) // pwwkew
	for i, v := range input {
		if _, exists := checker[v]; !exists {
			checker[v] = i
		} else {
			left = maximum(left, checker[v]+1)
			checker[v] = i
		}
		ans = maximum(ans, i-left+1)
	}
	return ans
}

func maximum(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}

// pwwkew
// p => p:0 left = 0 and ans = max(0, 0-0+1) => 1
// w =>  w:1, left = 0 & and = max(1, 0-1+1) = > 1
// w => ans =1 left = 0, => left = max(0, 2) => left =2, w:2, ans = max(1, 2-1+1) = 2
//
