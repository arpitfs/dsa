package round2

import "fmt"

func waterTrap() {
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
		rightMax[i] = maximum(rightMax[i+1], input[i])
	}

	for i := 0; i < len(input); i++ {
		water += manimum(leftMax[i], rightMax[i]) - input[i]
	}

	fmt.Println(water)
}

func maximum(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func manimum(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}
