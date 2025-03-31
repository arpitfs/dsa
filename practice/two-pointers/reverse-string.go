package twopointers

import "fmt"

func ReverseString() {
	input := "hello"
	convertedString := []byte(input)

	left := 0
	right := len(input) - 1

	for left < right {
		convertedString[left], convertedString[right] = convertedString[right], convertedString[left]
		left++
		right--
	}

	fmt.Println(string(convertedString))
}
