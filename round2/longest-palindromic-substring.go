package round2

import "fmt"

func longestPalindromSubString() {
	input := "babad"
	fmt.Println(getLongestPalindromicString(input))
}

func getLongestPalindromicString(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := findAllSubStrings(s, i, i)
		len2 := findAllSubStrings(s, i, i+1)
		maxLength := maximum(len1, len2)

		if maxLength > end-start {
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}
	}

	return s[start : end+1]
}

func findAllSubStrings(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return j - i - 1
}
