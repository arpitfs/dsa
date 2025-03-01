package string

import "fmt"

// Given a string, return all possible substrings.

func generatePossibleSubstrings() {
	input := "abc"
	result := []string{}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ {
			result = append(result, input[i:j])

		}
	}

	fmt.Println(result)
}
