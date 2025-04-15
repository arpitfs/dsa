package round2

import "fmt"

func rains() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	water := 0
	leftMax := make([]int, len(input))
	rightMax := make([]int, len(input))

	leftMax[0] = input[0]
	rightMax[len(input)-1] = input[len(input)-1]
	for i := 1; i < len(input); i++ {
		leftMax[i] = maximum(input[i], leftMax[i-1])
	}

	for i := len(input) - 2; i >= 0; i-- {
		rightMax[i] = maximum(input[i], rightMax[i+1])
	}

	for i := 0; i < len(input)-1; i++ {
		water += minimum(leftMax[i], rightMax[i]) - input[i]
	}

	fmt.Println(water)
}
