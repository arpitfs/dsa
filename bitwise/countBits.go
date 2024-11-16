package bitwise

import "fmt"

func countBits() {
	input := 9
	count := 0
	for input > 0 {
		lastBit := input & 1
		if lastBit == 1 {
			count++
		}
		input = input >> 1
	}

	fmt.Println(count)
}
