package twopointers

import "fmt"

//Remove duplicates in-place and return the new length.

func RemoveDuplicates() {
	input := []int{1, 1, 2, 2, 3, 4, 4}
	set := make(map[int]bool)
	result := []int{}
	left := 0
	right := len(input) - 1
	for left < right {
		if _, e := set[input[left]]; !e {
			result = append(result, input[left])
			set[input[left]] = true
		}

		if _, e := set[input[right]]; !e {
			result = append(result, input[right])
			set[input[right]] = true
		}

		left++
		right--
	}
	fmt.Println(len(result), result)
}
