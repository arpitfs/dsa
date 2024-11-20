package twopointer

import "fmt"

func getHouses() {
	input := []int{1, 3, 2, 1, 4, 1, 3, 2, 1, 1, 2}
	k := 8

	start := 0
	end := 0
	result := 0

	for end < len(input) {
		result = result + input[end]
		end++

		for result > k && start < end {
			result = result - input[start]
			start++
		}

		if result == k {
			fmt.Println(start, end-1)
		}

	}
}
