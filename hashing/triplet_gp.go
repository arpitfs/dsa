package hashing

import "fmt"

func countTriplet() {
	input := []int{1, 4, 10, 16, 4, 64, 1, 0}
	r := 4

	leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	for _, val := range input {
		leftMap[val] = 0
		rightMap[val] = rightMap[val] + 1
	}

	ans := 0

	for _, val := range input {
		rightMap[val]--
		if val%r == 0 {
			c := val * r
			a := val / r

			ans += leftMap[a] * rightMap[c]
		}

		leftMap[val]++
	}

	fmt.Println(ans)
}
