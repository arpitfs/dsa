package round2

import (
	"fmt"
	"strings"
)

func palindrome() {
	input := "raecear"
	strings.ToLower(input)
	var result string
	for _, v := range input {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			result += string(v)
		}
	}

	left := 0
	right := len(result) - 1
	for left < right {
		if result[left] != result[right] {
			fmt.Println("NO")
			break
		}
		left++
		right--
	}
}
