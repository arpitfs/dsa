package basic

import "fmt"

func maxSubarraysum() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	current := input[0]
	maximum := input[0]

	for i := 1; i < len(input); i++ {
		current = max(input[i], current+input[i])
		maximum = max(maximum, current)
	}

	fmt.Println(maximum)

}
