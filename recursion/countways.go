package recursion

import "fmt"

func countWays() {
	steps := 4
	// ways 1,2,3
	fmt.Println(count(steps))
}

func count(steps int) int {
	if steps < 0 {
		return 0
	}
	if steps == 0 {
		return 1
	}

	return count(steps-1) + count(steps-2) + count(steps-3)
}
