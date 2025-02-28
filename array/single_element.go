package array

import "fmt"

// Find the Single Element in a Duplicated Array

func findSingleElement() {
	input := []int{4, 1, 2, 3, 1, 2, 3}
	set := make(map[int]int)
	result := 0
	for _, val := range input {
		set[val]++
	}
	for k, val := range set {

		if val == 1 {
			result = k
		}
	}

	fmt.Println(result)
}
