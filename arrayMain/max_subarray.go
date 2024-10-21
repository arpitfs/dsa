package arrayMain

import "fmt"

func maxSubArraySum() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	current := input[0]
	max := input[0]
	for i := 0; i < len(input); i++ {
		current = maxValue(input[i], current+input[i])
		max = maxValue(max, current)
	}

	fmt.Println(max)
}

func maxValue(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}
