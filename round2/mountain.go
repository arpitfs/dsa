package round2

import "fmt"

func mountain() {
	input := []int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4}
	ans := 0

	for i := 1; i < len(input)-2; i++ {
		if input[i] > input[i-1] && input[i] > input[i+1] {
			total := 0
			j := i
			for j > 0 && input[j-1] < input[j] {
				total++
				j--
			}

			for i < len(input)-1 && input[i] > input[i+1] {
				total++
				i++
			}
			if total > ans {
				ans = total
			}
		}
	}

	fmt.Println(ans + 1)
}
