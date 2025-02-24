package basic

import "fmt"

func findNRC() {
	input := "leetcode"
	fmt.Printf("%c", firstNonRepeatingChar(input))
}

func firstNonRepeatingChar(s string) rune {
	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}

	for _, char := range s {
		if freq[char] == 1 {
			return char

		}
	}

	return '-'

}
