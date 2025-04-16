package round2

import "fmt"

func uniqueWindow() {
	input := "abcaeba"

	start := 0
	maxLen := 0
	startIndex := 0
	checker := make(map[rune]int)

	for i, ch := range input {
		if previousIndex, e := checker[ch]; e && previousIndex >= start {
			start = previousIndex + 1
		}

		checker[ch] = i
		if i-start+1 >= maxLen {
			maxLen = i - start + 1
			startIndex = start
		}

	}

	fmt.Println(input[startIndex : startIndex+maxLen])
}
