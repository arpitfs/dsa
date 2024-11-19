package binarysearch

import "fmt"

func frequency() {
	input := []int{1, 2, 2, 2, 3, 4, 5}
	number := 2

	lowerIndex := lowerBound(input, number)
	upperIndex := upperBound(input, number)

	fmt.Println("Frequency: ", upperIndex-lowerIndex+1)

}

func lowerBound(input []int, number int) int {
	start := 0
	end := len(input) - 1
	ans := -1
	for start <= end {
		mid := (start + end) / 2
		if input[mid] == number {
			ans = mid
			end = mid - 1
		} else if input[mid] > number {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}

func upperBound(input []int, number int) int {
	start := 0
	end := len(input) - 1
	ans := -1
	for start <= end {
		mid := (start + end) / 2
		if input[mid] == number {
			ans = mid
			start = mid + 1
		} else if input[mid] > number {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}
