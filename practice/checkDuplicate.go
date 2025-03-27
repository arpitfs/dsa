package practice

import "fmt"

func checkDuplicatePractice() {
	input := []int{2, 1, 5, 4, 5}
	fmt.Println(checkDuplicate(input))
}

func checkDuplicate(input []int) bool {
	set := make(map[int]int)
	for i := 0; i < len(input); i++ {
		if _, exists := set[input[i]]; !exists {
			set[input[i]] = i
		} else {
			return true
		}
	}
	return false
}
