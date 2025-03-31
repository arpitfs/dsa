package twopointers

import "fmt"

//Move all zeros to the end, maintaining the order.
func MoveAllZeros() {
	input := []int{0, 1, 0, 3, 12}
	length := len(input)
	result := make([]int, length)
	j := 0
	for _, v := range input {
		if v != 0 {
			result[j] = v
			j++
		}
	}

	fmt.Println(result)
}
