package basic

import "fmt"

func secondLargest() {
	input := []int{10, 45, 67, 89, 23}
	first := input[0]
	second := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > first && input[i] > second {
			second = first
			first = input[i]
		}
	}

	fmt.Println("Second Largest", second)
}
