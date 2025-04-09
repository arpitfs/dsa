package round2

import "fmt"

func findMaxInSlindingWindow() {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := []int{}
	for i := 0; i < len(input)-2; i++ {
		j := i
		currentMax := input[j]
		for j < i+k {
			if input[j] > currentMax {
				currentMax = input[j]
			}
			j++
		}
		result = append(result, currentMax)
	}
	fmt.Println(result)
}
