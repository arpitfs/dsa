package string

import (
	"fmt"
	"strings"
)

func palindrome() {
	input := "A man, a plan, canal, Panama!"
	input = "racecar"
	result := ""
	convertedInput := strings.ToLower(input)

	for _, val := range convertedInput {
		if (val >= 'a' && val <= 'z') || (val >= '0' && val <= '9') {
			result += string(val)
		}
	}
	fmt.Println(result)

	fmt.Println(checkPalindromeIn(result))

}

func checkPalindromeIn(result string) bool {
	left := 0
	right := len(result) - 1
	for left < right {
		if result[left] != result[right] {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}
