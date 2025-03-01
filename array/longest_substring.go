package array

import "fmt"

// Find the longest palindromic substring.
func lwrc() {

	input := "abba"
	left := 0
	ans := 0
	tracker := make(map[rune]int)

	for right := 0; right < len(input); right++ {
		if _, e := tracker[rune(input[right])]; !e {
			tracker[rune(input[right])] = right
		} else {
			left = max(left, tracker[rune(input[right])]+1)
			tracker[rune(input[right])] = right
		}

		ans = max(ans, right-left+1)
	}

	fmt.Println(ans)

}
