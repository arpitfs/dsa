package round2

import (
	"fmt"
	"math"
)

func isOutOfOrder(arr []int, i int) bool {
	n := len(arr)
	if i == 0 {
		return arr[i] > arr[i+1]
	}
	if i == n-1 {
		return arr[i] < arr[i-1]
	}
	return arr[i] > arr[i+1] || arr[i] < arr[i-1]
}

func subarraySort() {
	input := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}

	minOutOfOrder := math.MaxInt64
	maxOutOfOrder := math.MinInt64
	for i := 0; i < len(input); i++ {
		if isOutOfOrder(input, i) {
			if input[i] < minOutOfOrder {
				minOutOfOrder = input[i]
			}
			if input[i] > maxOutOfOrder {
				maxOutOfOrder = input[i]
			}
		}
	}

	left := 0
	for left < len(input) && input[left] <= minOutOfOrder {
		left++
	}

	right := len(input) - 1
	for right >= 0 && input[right] >= maxOutOfOrder {
		right--
	}
	fmt.Println(left, right)
}
