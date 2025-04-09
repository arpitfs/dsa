package round2

import "fmt"

type Output struct {
	f int
	s int
}

func twoSum() {
	input := []int{2, 7, 11, 15}
	target := 9
	output := []Output{}

	left := 0
	right := len(input) - 1
	for left < right {
		currentSum := input[left] + input[right]
		if currentSum == target {
			output = append(output, Output{f: left, s: right})
			left++
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	fmt.Println(output)

}
