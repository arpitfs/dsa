package twopointers

import "fmt"

// Find the total amount of water trapped.
func TrappingWater() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	n := len(input)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	water := 0

	leftMax[0] = input[0]
	for i := 1; i < n; i++ {
		leftMax[i] = maximumHeight(leftMax[i-1], input[i])
	}

	rightMax[n-1] = input[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = maximumHeight(rightMax[i+1], input[i])
	}

	for i := 0; i < n; i++ {
		water += mainimumHeight(leftMax[i], rightMax[i]) - input[i]
	}

	fmt.Println(water)
}

func maximumHeight(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func mainimumHeight(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}
