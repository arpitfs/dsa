package practice

import "fmt"

//Find the median of two sorted arrays.
func MedianSortedArray() {
	input1 := []int{1, 2, 3, 4}
	input2 := []int{2, 4, 5, 6}
	m, n := len(input1), len(input2)
	i, j := 0, 0
	mergedArray := make([]int, 0, m+n)

	for i < m && j < n {
		if input1[i] < input2[j] {
			mergedArray = append(mergedArray, input1[i])
			i++
		} else {
			mergedArray = append(mergedArray, input2[j])
			j++
		}
	}

	for i < m {
		mergedArray = append(mergedArray, input1[i])
		i++
	}

	for j < n {
		mergedArray = append(mergedArray, input2[j])
		j++
	}
	fmt.Println(mergedArray)

	total := m + n
	if total%2 == 1 {
		fmt.Println(float64(mergedArray[total/2]))
	} else {
		fmt.Println(float64(mergedArray[total/2-1]+mergedArray[total/2]) / 2.0)
	}
}
