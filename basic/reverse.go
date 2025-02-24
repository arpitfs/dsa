package basic

import "fmt"

func reverseString() {
	input := "hello"
	result := ""

	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}
	fmt.Println(result)
}
