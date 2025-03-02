package string

import (
	"fmt"
	"sort"
	"strings"
)

// Find all anagrams of a string in a list of strings.
func findAllAnagrams() {
	input := "listen"
	words := []string{"enlist", "google", "inlets", "banana", "silent", "tinsel"}
	fmt.Println(testAnangram(input, words))
}

func testAnangram(input string, words []string) []string {
	result := []string{}
	sortedInput := sortedString(input)
	for _, word := range words {
		sortedWord := sortedString(word)
		if sortedWord == sortedInput {
			result = append(result, word)
		}
	}
	return result
}

func sortedString(target string) string {
	slice := strings.Split(target, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
