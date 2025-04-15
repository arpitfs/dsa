package round2

import (
	"fmt"
	"sort"
)

type ResultTriplet struct {
	value1 int
	value2 int
	value3 int
}

func triplets() {
	input := []int{-1, 0, 1, 2, -1, -4}
	var result []ResultTriplet
	target := 0

	sort.Ints(input)

	for i := 0; i < len(input)-3; i++ {
		currentValue := input[i]
		leftValue := target - currentValue
		left := i + 1
		right := len(input) - 1
		for left < right {
			insideCurrentValue := input[left] + input[right]
			if insideCurrentValue == leftValue {
				newResult := ResultTriplet{value1: i, value2: left, value3: right}
				result = append(result, newResult)

				for left < len(input)-1 && input[left] == input[left+1] {
					left++
				}

				for right > 0 && input[right] == input[right-1] {
					right--
				}
				left++
				right--
			} else if insideCurrentValue < leftValue {
				left++
			} else {
				right--
			}
		}
	}

	fmt.Println(result)
}
