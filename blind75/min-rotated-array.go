package blind75

import "fmt"

// it should be O(log n)
// func minRotatedArray() {
// 	input := []int{4, 5, 6, 7, 0, 1, 2}

// 	for i := 0; i < len(input)-1; i++ {
// 		if input[i] > input[i+1] {
// 			fmt.Println(input[i+1])
// 			break
// 		}
// 	}

// }

func minRotatedArray() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(minRotatedArrayResult(input))
}

func minRotatedArrayResult(input []int) int {
	if len(input) == 1 {
		return input[0]
	}

	if len(input) == 2 {
		return input[0]
	}

	if input[0] < input[len(input)-1] {
		return input[0]
	}

	l := 0
	r := len(input) - 1

	for l < r {
		mid := l + (r-l)/2
		if input[mid+1] < input[mid] {
			return input[mid+1]
		}

		if input[mid] < input[mid-1] {
			return input[mid]
		}

		if input[l] < input[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return 0
}
