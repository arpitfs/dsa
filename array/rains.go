package array

import "fmt"

func rain() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	var leftArray [12]int
	var rightArray [12]int
	leftMax := 0
	rightMax := 0
	total := 0
	for i, value := range input {
		if value > leftMax {
			leftArray[i] = value
			leftMax = value
		} else {
			leftArray[i] = leftMax
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] > rightMax {
			rightArray[i] = input[i]
			rightMax = input[i]
		} else {
			rightArray[i] = rightMax
		}
	}

	for i := 0; i < len(input); i++ {
		minValue := min(leftArray[i], rightArray[i])
		total += minValue - input[i]
	}
	fmt.Println("Rain:= ", total)
}
