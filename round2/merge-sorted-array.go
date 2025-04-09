package round2

import "fmt"

func mergeSortedArray() {
	input1 := []int{1, 2, 3, 4}
	input2 := []int{2, 4, 5, 6}

	m, n := len(input1), len(input2)
	i, j := 0, 0
	result := make([]int, 0, m+n)

	for i < m && j < n {
		if input1[i] < input2[j] {
			result = append(result, input1[i])
			i++
		} else {
			result = append(result, input2[j])
			j++
		}
	}

	for i < m {
		result = append(result, input1[i])
		i++
	}

	for j < n {
		result = append(result, input2[j])
		j++
	}

	fmt.Println(result)
}
