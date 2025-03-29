package practice

import "fmt"

func mostWater() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	left := 0
	right := len(input) - 1
	water := 0
	for left < right {
		minHeight := min(input[left], input[right])
		l := right - left
		cw := minHeight * l
		if cw > water {
			water = cw
		}

		if input[right] > input[left] {
			left++
		} else {
			right--
		}
	}

	fmt.Println(water)
}

func min(value1, value2 int) int {
	if value1 < value2 {
		return value1
	}
	return value2
}
