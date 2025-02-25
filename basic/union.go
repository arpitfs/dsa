package basic

import "fmt"

func findUnion() {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 3, 4}
	result := []int{}
	set := make(map[int]bool)

	for _, val := range input1 {
		if !set[val] {
			set[val] = true
			result = append(result, val)
		}
	}

	for _, val := range input2 {
		if !set[val] {
			set[val] = true
			result = append(result, val)
		}
	}
	fmt.Println(result)
}
