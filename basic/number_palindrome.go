package basic

import "fmt"

func numberPalindrome() {
	input := 122
	original := input
	result := 0
	for input > 0 {
		lastDigit := input % 10
		result = result*10 + lastDigit
		input = input / 10
	}

	fmt.Println(result)
	if result == original {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
