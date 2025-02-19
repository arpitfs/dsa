package array

import "fmt"

func longestBand() {
	input := []int{1, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6}
	set := make(map[int]int)
	ans := 1
	for _, value := range input {
		set[value] = value
	}

	for i := 0; i < len(input); i++ {
		currentValue := set[input[i]]
		if _, exists := set[currentValue-1]; !exists {
			checkFlag := true
			count := 1
			data := 1
			for checkFlag {
				nextValue := input[i] + data
				if _, exists := set[nextValue]; exists {
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
