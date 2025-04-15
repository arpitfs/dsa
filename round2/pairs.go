package round2

import "fmt"

type ResultTwoSum struct {
	value1 int
	value2 int
}

func pairs() {
	input := []int{2, 7, 11, 15}
	target := 9
	var result []ResultTwoSum
	left := 0
	right := len(input) - 1

	for left < right {
		currentValue := input[left] + input[right]
		if currentValue == target {
			newResult := ResultTwoSum{value1: left, value2: right}
			result = append(result, newResult)
			left++
		} else if currentValue < target {
			left++
		} else {
			right--
		}
	}

	fmt.Println(result)

}
