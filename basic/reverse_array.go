package basic

import "fmt"

func reverse() {
	arr1 := []int{1, 2, 3, 4, 5}

	start := 0
	end := len(arr1) - 1

	for start < end {
		arr1[start], arr1[end] = arr1[end], arr1[start]
		start++
		end--
	}
	fmt.Println(arr1)
}
