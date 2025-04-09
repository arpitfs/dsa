package round2

import (
	"fmt"
	"sort"
)

type OutputThreeSum struct {
	f int
	s int
	t int
}

func threeSum() {
	input := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	sort.Ints(input)
	result := []OutputThreeSum{}

	for i := 0; i < len(input)-2; i++ {
		currentValue := input[i]
		left := i + 1
		right := len(input) - 1
		findTarget := target - currentValue

		for left < right {
			internalSum := input[left] + input[right]
			if internalSum == findTarget {
				result = append(result, OutputThreeSum{f: i, s: left, t: right})
				for left < len(input)-1 && input[left] == input[left+1] {
					left++
				}

				for right > 0 && input[right] == input[right-1] {
					right--
				}

				left++
				right--

			} else if internalSum < findTarget {
				left++

			} else {
				right--
			}
		}
	}

	fmt.Println(result)

}
