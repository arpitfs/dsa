package round2

import (
	"fmt"
	"math"
)

func minWindowSubString() {
	input := "ADOBECODEBANC"
	t := "ABC"

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, have := 0, 0, 0
	needLength := len(need)
	window := make(map[byte]int)
	result := []int{-1, -1}
	resultLength := math.MaxInt32

	for right < len(input) {
		c := input[right]
		window[c]++

		if _, e := need[c]; e && window[c] == need[c] {
			have++
		}

		for have == needLength {
			if (right - left + 1) < resultLength {
				result[0] = left
				result[1] = right
				resultLength = right - left + 1
			}

			window[input[left]]--
			if _, e := need[input[left]]; e && window[input[left]] < need[input[left]] {
				have--
			}
			left++
		}
		right++
	}

	if resultLength == math.MaxInt32 {
		fmt.Println("")
	}

	fmt.Println(input[result[0] : result[1]+1])
}
