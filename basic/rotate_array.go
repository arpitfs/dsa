package basic

import "fmt"

func rotateArray() {
	input := []int{1, 2, 3, 4, 5}
	k := 2

	resultArray := make([]int, 5)
	z := 0
	for i := k + 1; i < len(input); i++ {
		resultArray[z] = input[i]
		z++
	}

	len := 2

	for i := 0; i <= k; i++ {
		resultArray[len] = input[i]
		len++
	}
	fmt.Println(resultArray)
}
