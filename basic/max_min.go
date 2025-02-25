package basic

import "fmt"

func maxMin() {
	input := []int{3, 5, 1, 2, 4, 8}
	max := input[0]
	min := input[0]

	for i := 1; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}

		if input[i] < min {
			min = input[i]
		}
	}

	fmt.Println("Max:Min", max, ":", min)
}
