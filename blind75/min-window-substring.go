	package blind75

import (
	"fmt"
	"math"
)

// Given two strings s and t of lengths m and n respectively,
// return the minimum window substring of s such that every character in t (including duplicates) is
// included in the window. If there is no such substring, return the empty string ""

func minWindow() {
	input := "ADOBECODEBANC"
	subString := "ABC"
	fmt.Println(minWindowSubstring(input, subString))

}

func minWindowSubstring(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, have := 0, 0, 0
	needCount := len(need)
	window := make(map[byte]int)
	result := []int{-1, -1}
	resultLength := math.MaxInt32

	for right < len(s) {
		c := s[right]
		window[c]++

		if _, e := need[c]; e && window[c] == need[c] {
			have++
		}

		for have == needCount {
			if (right - left + 1) < resultLength {
				result[0] = left
				result[1] = right
				resultLength = right - left + 1
			}

			window[s[left]]--
			if _, ok := need[s[left]]; ok && window[s[left]] < need[s[left]] {
				have--
			}
			left++
		}
		right++
	}

	if resultLength == math.MaxInt32 {
		return ""
	}

	return s[result[0] : result[1]+1]
}
