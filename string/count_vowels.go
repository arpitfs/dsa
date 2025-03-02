package string

import (
	"fmt"
	"strings"
)

// Given a string, count the number of vowels.
func countVowels() {
	input := "abcedef"
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	vowelsString := ""
	for _, v := range vowels {
		vowelsString += string(v)

	}

	count := 0
	for _, val := range input {
		value := string(val)
		if strings.Contains(vowelsString, value) {
			count++
		}
	}

	fmt.Println(count)
}
