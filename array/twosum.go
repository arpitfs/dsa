package array

import "fmt"

func twoSum() {
	input := []int{2, 7, 11, 15}
	target := 9

	set := make(map[int]int)

	for i := 0; i <= len(input)-1; i++ {
		data := target - input[i]
		if _, exists := set[data]; exists {
			fmt.Println(i, set[data])
		} else {
			set[input[i]] = i
		}

	}
}
