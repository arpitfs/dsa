package recursion

import "fmt"

func subsetSequence() {
	input := []int{1, 2, 3, 4, 5}
	sum := 6

	fmt.Println(subset(input, sum, 0))
}

func subset(input []int, sum, i int) int {
	if i == len(input) {
		if sum == 0 {
			return 1
		}
		return 0
	}

	inc := subset(input, sum-input[i], i+1)
	exc := subset(input, sum, i+1)

	return inc + exc
}
