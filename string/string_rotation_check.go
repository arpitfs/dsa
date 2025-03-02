package string

import (
	"fmt"
	"strings"
)

// Check if a string is a rotation of another string.

func checkRotationString() {
	input := "waterbottle"
	rotatedInput := "erbottlewat"

	fmt.Println(checkRotation(input, rotatedInput))
	fmt.Println(checkManual(input, rotatedInput))

}

func checkRotation(input, rotatedInput string) bool {
	if len(input) != len(rotatedInput) {
		return false
	}

	return strings.Contains(input+input, rotatedInput)
}

func checkManual(input, rotatedInput string) bool {
	if len(input) != len(rotatedInput) {
		return false
	}
	//fmt.Println(input[3:], input[:3])
	//input := "waterbottle"
	//rotatedInput := "erbottlewat"
	for i := 0; i < len(input); i++ {

		if input[i:]+input[:i] == rotatedInput {
			return true
		}
	}

	return false
}
