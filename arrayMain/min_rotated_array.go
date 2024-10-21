package arrayMain

import "fmt"

func minRotatedArray() {
	input := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	ans := findAns(input)
	fmt.Println(ans)
}

func findAns(input []int) int {
	if len(input) == 1 {
		return input[0]
	}

	if len(input) == 2 {
		return input[1]
	}

	if input[0] < input[len(input)-1] {
		return input[0]
	}

	left := 0
	right := len(input) - 1
	for left <= right {
		mid := left + (right-left)/2

		// if mid + 1 is decreasing
		if input[mid+1] < input[mid] {
			return input[mid+1]
		}
		// if mid is decreasing

		if mid < input[mid-1] {
			return input[mid]
		}
		// discard the sorted
		if input[left] < input[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0

}
