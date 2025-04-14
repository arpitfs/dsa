package round2

import "fmt"

func jumpGame() {
	input := []int{2, 3, 1, 1, 4}
	maxJumps := 0

	for i, v := range input {
		if i >= maxJumps {
			fmt.Println("False")
			break
		}

		if i+v > maxJumps {
			maxJumps = i + v
		}
	}
	fmt.Println("True")
}
