package practice

import (
	"fmt"
	"sort"
)

type ThreeSum struct {
	f int
	s int
	t int
}

func threeSumPractice() {
	input := []int{-1, 0, 1, 2, -1, -4}
	target := 0

	sort.Ints(input)
	var result []ThreeSum
	// -4, -1,-1,0,1,2
	for i := 0; i < len(input)-3; i++ {
		if i == 0 || input[i] != input[i-1] {
			currentValue := input[i]
			left := i + 1
			right := len(input) - 1
			currentSum := target - currentValue
			for left < right {
				if input[left]+input[right] == currentSum {
					r := ThreeSum{f: i, s: left, t: right}
					result = append(result, r)

					for left < len(input)-1 && input[left] == input[left+1] {
						left++
					}
					for right > 0 && input[right] == input[right-1] {
						right--
					}
					left++
					right--
				} else if input[left]+input[right] < target {
					left++
				} else {
					right--
				}
			}
		}

	}
	fmt.Println(result)
}
