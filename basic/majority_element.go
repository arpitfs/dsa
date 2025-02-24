package basic

import "fmt"

func majorityElement() {
	input := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	set := make(map[int]int)

	for _, val := range input {
		if _, exists := set[val]; !exists {
			set[val] = 1
		} else {
			set[val] = set[val] + 1
		}
	}

	result := 0
	for _, val := range set {
		if val > result {
			result = val
		}

	}

	fmt.Println(result)
}
