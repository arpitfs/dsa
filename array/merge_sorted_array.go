package array

import "fmt"

// Merge Two Sorted Arrays Without Extra Space
func mergeSortedArray() {
	input1 := []int{1, 3, 5}
	input2 := []int{2, 4, 6}

	fmt.Println(mergeArray(input1, input2))

}

func mergeArray(input1, input2 []int) []int {

	result := []int{}
	i, j := 0, 0

	for i < len(input1) && j < len(input2) {
		if input1[i] < input2[j] {
			result = append(result, input1[i])
			i++
		} else {
			result = append(result, input2[j])
			j++
		}
	}

	for i < len(input1) {
		result = append(result, input1[i])
		i++
	}

	for j < len(input2) {
		result = append(result, input2[j])
		j++
	}

	return result

}
