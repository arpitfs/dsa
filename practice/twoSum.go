package practice

import "fmt"

type Result struct {
	first  int
	second int
}

func twoSumPractice() {
	input := []int{2, 7, 11, 15}
	target := 9
	var result []Result
	left := 0
	right := len(input) - 1

	for left < right {
		currentSum := input[left] + input[right]
		if currentSum == target {
			currentResult := Result{first: left, second: right}
			result = append(result, currentResult)
			left++
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	fmt.Println(result)
}
