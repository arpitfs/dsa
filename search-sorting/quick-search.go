package searchsorting

import "fmt"

func quicksearch() {
	input := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println(quickSelect(input, 0, len(input)-1, 3))
}

func quickSelect(input []int, i, j, k int) int {
	p := partition(input, i, j)
	if p == k {
		return input[p]
	} else if k < p {
		return quickSelect(input, i, p-1, k)
	} else {
		return quickSelect(input, p+1, j, k)
	}

}
