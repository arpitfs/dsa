package searchsorting

import "fmt"

func quicksort() {
	input := []int{38, 27, 43, 3, 9, 82, 10}
	doQuickSort(input, 0, len(input)-1)
	fmt.Println("Sorted Using Quick Sort:", input)
}

func doQuickSort(input []int, i, j int) {
	if i < j {
		p := partition(input, i, j)
		doQuickSort(input, i, p-1)
		doQuickSort(input, p+1, j)
	}
}

func partition(input []int, i, j int) int {
	pivot := input[j]
	k := i - 1
	for m := i; m < j; m++ {
		if input[m] < pivot {
			k++
			input[m], input[k] = input[k], input[m]
		}
	}

	input[k+1], input[j] = input[j], input[k+1]

	return k + 1
}
