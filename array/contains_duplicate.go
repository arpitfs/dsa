package array

import "fmt"

func containsDuplicate() {
	input := []int{2, 1, 5, 4, 5}
	set := make(map[int]int)

	for i := 0; i <= len(input)-1; i++ {
		if _, exists := set[input[i]]; exists {
			fmt.Println("True")
			break
		} else {
			set[input[i]] = input[i]
		}
	}
}
