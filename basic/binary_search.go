package basic

import "fmt"

func binarySearch() {

	input := []int{10, 23, 45, 67, 89}
	search := 89
	fmt.Println(bs(input, search))

}

func bs(input []int, search int) int {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (low + high) / 2
		if input[mid] == search {
			return mid
		} else if search > input[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
