package blind75

import "fmt"

func maxSubArraySum() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	currentSum := input[0]
	maxSum := input[0]
	for i := 0; i < len(input); i++ {
		currentSum = maximum(input[i], currentSum+input[i])
		maxSum = maximum(currentSum, maxSum)
	}
	fmt.Println(maxSum)
}

func maximum(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
