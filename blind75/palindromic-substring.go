package blind75

import "fmt"

func countSubstring() {
	input := "abc"
	fmt.Println(palindromicSubstrings(input))
}

func palindromicSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += findSubstring(s, i, i)
		count += findSubstring(s, i, i+1)
	}
	return count
}

func findSubstring(s string, i, j int) int {
	count := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		if s[i] == s[j] {
			count++
			i--
			j++
		}
	}
	return count
}
