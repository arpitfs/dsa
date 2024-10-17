// package main

// import "fmt"

// func main() {
// 	input := "abba"
// 	fmt.Println(longestSubstring(input))
// 	input = "pwwkew"
// 	fmt.Println(longestSubstring(input))
// }

// func longestSubstring(input string) int {
// 	ans := 0
// 	left := 0
// 	checker := make(map[rune]int)
// 	for i, v := range input {
// 		if _, exists := checker[v]; !exists {
// 			checker[v] = i
// 		} else {
// 			left = maximum(left, checker[v]+1)
// 			checker[v] = i
// 		}
// 		ans = maximum(ans, i-left+1)
// 	}
// 	return ans
// }

// func maximum(value1, value2 int) int {
// 	if value1 > value2 {
// 		return value1
// 	}
// 	return value2
// }
