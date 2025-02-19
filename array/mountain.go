package array

import "fmt"

func mountain() {
	input := []int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4}
	ans := 0
	for i := 1; i < len(input)-2; i++ {
		if input[i] > input[i+1] && input[i] > input[i-1] {
			count := 1
			j := i
			for j >= 1 && input[j] > input[j-1] {
				j--
				count++
			}

			for i <= len(input)-2 && input[i] > input[i+1] {
				i++
				count++
			}

			if count > ans {
				ans = count
			}

		}

	}
	fmt.Println(ans)
}
