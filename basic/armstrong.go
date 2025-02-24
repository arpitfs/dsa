package basic

import (
	"fmt"
	"math"
)

func armstrong() {

	input := 153
	lengthCheck := input
	original := input
	length := 0
	result := 0

	for lengthCheck > 0 {
		lengthCheck = lengthCheck / 10
		length++
	}

	for input > 0 {
		lastDigit := input % 10
		result += int(math.Pow(float64(lastDigit), float64(length)))
		input = input / 10
	}

	if original == result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
