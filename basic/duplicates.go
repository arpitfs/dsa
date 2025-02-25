package basic

import "fmt"

func findDuplicate() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := []int{}
	set := make(map[int]bool)

	for _, val := range input {
		if !set[val] {
			set[val] = true
		} else {
			result = append(result, val)
		}
	}

	fmt.Println(result)
}
