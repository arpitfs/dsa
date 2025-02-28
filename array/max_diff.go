package array

import (
	"fmt"
)

//  maxDifference finds the maximum difference between two elements such that the larger element appears after the smaller element

func maxDiff() {
	input := []int{7, 1, 5, 3, 6, 4}
	minElement := input[0]
	maxDifference := 0
	for i := 1; i < len(input); i++ {
		if input[i] < minElement {
			minElement = input[i]
		} else {
			maxDifference = max(maxDifference, input[i]-minElement)
		}
	}

	fmt.Println(maxDifference)
}
