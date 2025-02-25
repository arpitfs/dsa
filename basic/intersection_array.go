package basic

import "fmt"

func findIntersection() {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 3, 4}

	result := []int{}
	set := make(map[int]int)

	for _, val := range input1 {
		if _, e := set[val]; !e {
			set[val] = val
		}
	}

	for _, val := range input2 {
		if _, e := set[val]; e {
			result = append(result, set[val])
		}
	}

	fmt.Println(result)
}
