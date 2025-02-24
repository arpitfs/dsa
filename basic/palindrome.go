package basic

import (
	"fmt"
	"strings"
)

func checkpalindrome() {
	input := "A man, a plan, canal, Panama!"

	convertedInput := strings.ToLower(input)
	var result string
	for _, c := range convertedInput {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			result += string(c)
		}
	}

	fmt.Println(isPalin(result))

}

func isPalin(result string) bool {

	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-1-i] {
			return false
		}

	}
	return true
}
