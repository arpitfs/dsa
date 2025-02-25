package basic

import "fmt"

func arraySorted() {
	input := []int{2, 3, 5, 6}
	firstNumber := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] != firstNumber+1 {
			fmt.Println("No")
			break
		}
		firstNumber++
	}
}
