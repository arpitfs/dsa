package binarysearch

import "fmt"

func RotatedSearch() {
	input := []int{4, 5, 6, 7, 0, 1, 2, 3}
	search := 2
	start := 0
	end := len(input) - 1

	for start <= end {
		mid := (start + end) / 2
		if input[mid] == search {
			fmt.Println(mid)
		}

		if input[start] <= input[mid] {
			if search > input[start] && search <= input[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if search >= input[mid] && search <= input[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}

		}
	}
}
