package array

import "fmt"

// Find the First and Last Position of an Element in Sorted Array
type FirstAndLastResult struct {
	first int
	last  int
}

func firstAndLast() {
	input := []int{5, 7, 7, 8, 8, 10}
	target := 8

	fmt.Println(findFirst(input, target), findLast(input, target))

}

// func findIndexs(input []int, target int) FirstAndLastResult {
// 	if len(input) <= 0 {
// 		return FirstAndLastResult{first: -1, last: -1}
// 	}
// 	firstIndex, secondIndex := 0, 0
// 	firstFound := false
// 	for i := 0; i < len(input); i++ {
// 		if input[i] == target && !firstFound {
// 			firstIndex = i
// 			secondIndex = i
// 			firstFound = true
// 		} else if input[i] == target && firstFound && input[i+1] > target {
// 			secondIndex = i
// 		}
// 	}

// 	return FirstAndLastResult{first: firstIndex, last: secondIndex}
// }

// Using BinarySeach

func findFirst(input []int, target int) int {
	left, right := 0, len(input)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if input[mid] == target {
			result = mid
			right = mid - 1
		} else if input[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result

}

func findLast(input []int, target int) int {
	left, right := 0, len(input)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if input[mid] == target {
			result = mid
			left = mid + 1
		} else if input[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result

}
