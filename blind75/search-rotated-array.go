package blind75

import (
	"fmt"
)

// seach with O(log n)
func searchSortedArray() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 4

	fmt.Println(searchRotatedSortArray(input, target))
}
func searchRotatedSortArray(input []int, target int) int {

	l := 0
	r := len(input) - 1

	for l <= r {
		mid := l + (r-l)/2
		if input[mid] == target {
			return input[mid]
		}

		if input[l] < input[mid] {
			if input[l] <= target && target < input[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if input[mid] >= target && target > input[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
