package array

import (
	"fmt"
	"sort"
)

// Find the Longest Consecutive Sequence
func longestConsecutiveSequence() {
	input := []int{202, 203, 4, 201, 1, 7, 2}

	sort.Ints(input) //1,2,3,4,200,400

	longest, current := 1, 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			continue
		} else if input[i] == input[i-1]+1 {
			current++
		} else {
			longest = max(longest, current)
			current = 1
		}
	}

	fmt.Println(max(longest, current))

}
