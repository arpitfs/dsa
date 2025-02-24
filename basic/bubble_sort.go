package basic

import "fmt"

func bubbleSort() {
	input := []int{10, 45, 67, 89, 23}

	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}

	fmt.Println(input)
}
