package string

import "fmt"

// Implement a function to remove duplicate characters from a string.

func removeDuplicates() {
	input := "abacabad"
	set := make(map[rune]int)

	for _, val := range input {
		set[val]++
	}

	result := ""
	for _, val := range input {
		if set[val] != 0 {
			result += string(val)
			set[val] = 0
		}
	}

	fmt.Println(result)
}
