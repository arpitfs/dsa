package string

import "fmt"

// Implement a function to count the number of words in a sentence.

func countWords() {
	input := "The quick brown fox jumps over the lazy dog"
	count := 0
	for _, val := range input {
		if val == ' ' {
			count++
		}
	}

	fmt.Println(count + 1)
}
