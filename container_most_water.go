package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	left := 0
	right := len(input) - 1
	water := 0
	for left < right {
		l := right - left
		h := min(input[left], input[right])
		ans := l * h
		if ans > water {
			water = ans
		}

		if left > right {
			right--
		} else {
			left++
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
