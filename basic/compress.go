package basic

import (
	"fmt"
	"strconv"
)

func compress() {
	input := "aaabbbcccc"
	set := make(map[rune]int)
	for _, val := range input {
		set[val-rune('a')]++
	}
	result := ""
	fmt.Println(set)
	for i, v := range set {
		x := i + rune('a')
		result += string(x)
		result += strconv.Itoa(v)
	}
	fmt.Println(result)

}
