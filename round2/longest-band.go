package round2

import "fmt"

func longestBand() {
	input := []int{1, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6}

	checker := make(map[int]int)

	for i, v := range input {
		checker[v] = i
	}

	ans := 0

	for i := 0; i < len(input)-1; i++ {
		checkValue := input[i] - 1
		if _, e := checker[checkValue]; !e {
			checkFlag := true
			count := 1
			data := 1
			for checkFlag {
				nextValue := input[i] + data
				if _, exists := checker[nextValue]; exists {
					count++
					data++
				} else {
					if count > ans {
						ans = count
					}
					checkFlag = false
				}
			}
		}
	}
	fmt.Println(ans)
}
