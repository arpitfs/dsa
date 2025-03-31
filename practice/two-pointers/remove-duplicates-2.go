package twopointers

import "fmt"

// Remove Duplicates from Sorted Array II

//nums = [1, 1, 1, 2, 2, 3]
//[1, 1, 2, 2, 3]
//Length = 5

func RemoveDuplicates2() {
	input := []int{1, 1, 1, 2, 2, 3}
	set := make(map[int]int)
	result := []int{}
	for _, v := range input {
		if _, e := set[v]; !e {
			set[v] = 1
			result = append(result, v)
		} else {
			currentCount := set[v]
			if currentCount < 2 {
				result = append(result, v)
				set[v]++
			}
		}
	}

	fmt.Println(len(result), result)
}
