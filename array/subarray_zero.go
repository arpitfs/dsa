package array

import "fmt"

// Find if a Subarray with Sum 0 Exists

func subarrayZero() {
	input := []int{4, 2, -3, 1, 6}
	target := 0
	fmt.Println(checkSubArrayZero(input, target))

}

func checkSubArrayZero(input []int, target int) bool {
	set := make(map[int]bool)

	//set[0] = true

	for _, val := range input {
		target += val
		if set[target] {
			return true
		}
		set[target] = true
	}

	return false
}
