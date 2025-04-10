package round2

import "fmt"

func palindromicSubstring() {
	input := "abc"
	fmt.Println(countOfPalindromicSubstring(input))
}

func countOfPalindromicSubstring(s string) int {
	totalCount := 0
	for i, _ := range s {
		count1 := findAllSubStrings(s, i, i)
		count2 := findAllSubStrings(s, i, i+1)
		totalCount += count1
		totalCount += count2
	}
	return totalCount
}

func findSubStrings(s string, i, j int) int {
	count := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		count++
		i--
		j++
	}
	return count
}
