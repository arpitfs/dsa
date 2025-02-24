package basic

import "fmt"

type Result struct {
	value1 int
	value2 int
}

func findPairs() {
	input := []int{1, 2, 3, 4, 5, 6}
	target := 7
	set := make(map[int]int)
	var result []Result

	for i := 0; i < len(input); i++ {
		requriedValue := target - input[i]
		if _, exists := set[requriedValue]; exists {
			data := Result{value1: i, value2: set[requriedValue]}
			result = append(result, data)
		} else {
			set[input[i]] = i
		}
	}

	fmt.Println(result)

}
