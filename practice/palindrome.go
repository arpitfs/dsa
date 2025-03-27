package practice

import (
	"fmt"
	"strings"
)

func palindromePractice() {
	//input := "A man, a plan, canal, Panama!"
	input := "racecar"
	convertedString := strings.ToLower(input)
	var result string

	for _, c := range convertedString {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			result += string(c)
		}
	}

	// for i := 0; i < len(result)/2; i++ {
	// 	if result[i] != result[len(result)-1-i] {
	// 		fmt.Println(false)
	// 	}
	// }

	left := 0
	right := len(result) - 1
	for left < right {
		if result[left] != result[right] {
			fmt.Println(false)
		} else {
			left++
			right--
		}
	}

}
