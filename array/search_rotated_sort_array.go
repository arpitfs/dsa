package array

import "fmt"

func searchRotatedSortArray() {
	input := []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	target := 81
	ans := search(input, target)
	fmt.Println(ans)
}

func search(input []int, target int) int {
	if input[0] == target {
		return 0
	}

	if input[1] == target {
		return 1
	}

	left := 0
	right := len(input) - 1

	for left <= right {
		mid := left + (right-left)/2
		if input[mid] == target {
			return mid
		}
		if input[left] < input[mid] { // left to mid is sorted
			if input[left] <= target && target < input[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right to mid is sorted
			if input[mid] < target && target < input[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}
