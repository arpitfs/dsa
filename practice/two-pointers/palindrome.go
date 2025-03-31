package twopointers

import (
	"fmt"
	"strings"
)

func Palindrome() {
	input := "A man, a plana canal: Panama"
	convertedInput := strings.ToLower(input)
	var s strings.Builder
	for _, v := range convertedInput {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			s.WriteString(string(v))
		}
	}

	fmt.Println(checkPalindrome([]byte(s.String())))
}

func checkPalindrome(input []byte) bool {
	left := 0
	right := len(input) - 1

	for left < right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}

	return true
}
