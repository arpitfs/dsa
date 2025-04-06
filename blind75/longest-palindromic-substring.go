package blind75

import "fmt"

func longestPalindromSubString() {
	input := "babad"
	fmt.Println(getAllPalindromicString(input))
}

func getAllPalindromicString(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := findStringFromInitialString(s, i, i)
		len2 := findStringFromInitialString(s, i, i+1)
		maxLength := maximum(len1, len2)

		if maxLength > end-start {
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}

	}
	return s[start : end+1]
}

func findStringFromInitialString(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1
}
