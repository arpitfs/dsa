package basic

import "fmt"

func moveZeros() {
	input := []int{0, 1, 2, 0, 3, 0, 4}
	len := len(input)
	fmt.Println(len)
	result := make([]int, len)
	j := 0
	for i := 0; i < len; i++ {
		if input[i] != 0 {
			result[j] = input[i]
			j++
		}
	}

	fmt.Println(result)

}
