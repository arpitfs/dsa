package basic

import "fmt"

func sort012() {
	input := []int{2, 0, 2, 1, 1, 0}
	low, mid, high := 0, 0, len(input)-1

	for mid <= high {
		switch input[mid] {
		case 0:
			input[low], input[mid] = input[mid], input[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			input[mid], input[high] = input[high], input[mid]
			high--
		}
	}

	fmt.Println(input)
}
