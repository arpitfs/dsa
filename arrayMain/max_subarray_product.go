package arrayMain

import (
	"fmt"
)

func maxSubArrayProduct() {
	input := []int{-1, -2, -3, 0, 3, 5, -1, -2}
	min := 1
	max := 1
	ans := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			min = 1
			max = 1
			ans = maxOfThree(ans, input[i], 0)
			fmt.Println(i, min, max)
			continue
		}
		mi := min
		ma := max
		min = minOfThree(input[i], mi*input[i], ma*input[i])
		max = maxOfThree(input[i], mi*input[i], ma*input[i])
		fmt.Println(i, min, max)
		if max > ans {
			ans = max
		}

	}

	fmt.Println(ans)
}

func maxOfThree(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	return max
}

func minOfThree(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	return min
}
