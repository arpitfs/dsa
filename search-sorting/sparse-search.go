package searchsorting

import (
	"fmt"
	"strings"
)

func searchString() {
	input := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	target := "dad"
	result := SparseSearch(input, target, 0, len(input)-1)
	fmt.Println("Result: ", result)
}

func SparseSearch(input []string, target string, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	mid_left, mid_right := mid, mid

	for {
		if mid_left < left && mid_right > right {
			return -1
		}

		if mid_left >= left && input[mid_left] != "" {
			mid = mid_left
			break
		}
		if mid_right <= right && input[mid_right] != "" {
			mid = mid_right
			break
		}
		mid_left--
		mid_right++
	}

	if input[mid] == target {
		return mid
	} else if strings.Compare(input[mid], target) < 0 {
		return SparseSearch(input, target, mid+1, right)
	} else {
		return SparseSearch(input, target, left, mid-1)
	}
}
