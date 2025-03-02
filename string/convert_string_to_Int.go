package string

import (
	"fmt"
	"math"
	"unicode"
)

// Convert a string to an integer

func convertStringToInt() {
	input := "2135"
	sign := 1
	i := 0
	length := len(input)
	result := 0

	for i < length && unicode.IsSpace(rune(input[i])) {
		i++
	}

	if i < length && (input[i] == '-' || input[i] == '+') {
		if input[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < length && unicode.IsDigit(rune(input[i])) {
		digit := int(input[i] - '0')
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				fmt.Println("MaxInt")
				break
			}
		}
		result = result*10 + digit
		i++
	}

	fmt.Println(result * sign)

}

func convertSimpleIntToDigit() {
	input := "2135"
	result := 0
	for _, val := range input {
		digit := int(val - '0')
		result = result*10 + digit
	}

	fmt.Println(result)
}
