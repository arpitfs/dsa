package array

import "fmt"

func jump_game() {
	input := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 2, 5}
	isPossible, length := getMinJumps(input)
	fmt.Println("isPossible: ", isPossible, " Min Steps : ", length)
}

func getMinJumps(input []int) (bool, int) {
	reachable := 0
	maxSteps := 0
	for i := 1; i < len(input) && i <= reachable; i++ {
		maxSteps++
		reachable = maximum(reachable, i+input[i])
		if reachable >= len(input)-1 {
			return true, maxSteps
		}
	}
	return false, 0
}

func maximum(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}
