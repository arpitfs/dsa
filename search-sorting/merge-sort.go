package searchsorting

import "fmt"

func mergeSort() {
	input := []int{38, 27, 43, 3, 9, 82, 10}
	sortedArr := doMergeSort(input)
	fmt.Println("Sorted Array: ", sortedArr)
}

func doMergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	mid := len(input) / 2
	left := doMergeSort(input[:mid])
	right := doMergeSort(input[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
