package round2

import "fmt"

func countSubset() {
	input := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Println(countWays(input, len(input), 0, target))
}

func countWays(input []int, last, start, target int) int {
	if start == last {
		if target == 0 {
			return 1
		}
		return 0
	}

	inc := countWays(input, last, start+1, target-input[start])
	exc := countWays(input, last, start+1, target)
	return inc + exc
}
