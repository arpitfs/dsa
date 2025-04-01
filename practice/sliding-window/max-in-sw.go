package slidingwindow

import "fmt"

// Find the maximum in each sliding window of size k.
func MaxInSlidingWindow() {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	i := 0
	result := []int{}
	for i < len(input)-2 {
		j := i
		currentMax := 0
		for j < i+k {
			if input[j] > currentMax {
				currentMax = input[j]
			}
			j++
		}
		result = append(result, currentMax)
		i++
	}

	fmt.Println(result)
}
