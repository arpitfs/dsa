package twopointers

import "fmt"

// Find indices of two numbers that add up to the target.
func TargetSum() {
	input := []int{1, 2, 3, 4, 6}
	target := 6
	fmt.Println(checkSum(input, target))
}

func checkSum(input []int, target int) (int, int) {
	left := 0
	right := len(input) - 1
	for left < right {
		currentSum := input[left] + input[right]
		if currentSum == target {
			return left, right
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return -1, -1
}
