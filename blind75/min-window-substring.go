package blind75

import "fmt"

// Given two strings s and t of lengths m and n respectively,
// return the minimum window substring of s such that every character in t (including duplicates) is
// included in the window. If there is no such substring, return the empty string ""

func minWindow() {
	input := "ADOBECODEBANC"
	subString := "ABC"
	fmt.Println(input, subString)

}

// func getMinimumWindwo(input, subString string) string {
// 	if len(input) == 0 || len(subString) == 0 {
// 		return ""
// 	}

// 	tracker := make(map[byte]int)
// 	for i := 0; i < len(subString); i++ {
// 		tracker[subString[i]]++
// 	}

// 	l, r, minLeft, minSize, count := 0, 0, 0, math.MaxInt32, 0
// 	windowFreq := make(map[byte]int)
// 	for r < len(input) {
// 		c := subString[r]
// 		if tracker[c] > 0 {
// 			windowFreq[c]++
// 			if windowFreq[c] <= tracker[c] {
// 				count++
// 			}
// 		}
// 		r++
// 		if r-l < minSize {
// 			minSize = r - l
// 			minLeft = l
// 		}
// 	}
// }
