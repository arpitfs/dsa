package basic

import "fmt"

// func missing_number() {
// 	input := []int{1, 2, 3}
// 	n := 4
// 	missingNumber := 0
// 	for i := 0; i < len(input); i++ {
// 		if input[i] != i+1 {
// 			fmt.Println("Missing Number: ", i+1)
// 			missingNumber = i + 1
// 			break
// 		}
// 	}
// 	if missingNumber == 0 {
// 		fmt.Println("Missing Number: ", input[n-2]+1)
// 	}
// }

func missing_number() {
	input := []int{1, 2, 4, 5, 6}
	n := 6
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, val := range input {
		actualSum += val
	}
	fmt.Println("Missing Number: ", expectedSum-actualSum)
}
