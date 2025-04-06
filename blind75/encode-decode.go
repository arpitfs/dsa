package blind75

import (
	"fmt"
	"strconv"
	"strings"
)

func encodeDecode() {
	input := []string{"hello", "world", "go"}
	result := encode(input)
	list := decode(result)
	fmt.Println(list)
}

func encode(input []string) string {
	var encoded strings.Builder
	for _, str := range input {
		encoded.WriteString(strconv.Itoa(len(str)))
		encoded.WriteString("#")
		encoded.WriteString(str)
	}
	return encoded.String()
}

func decode(s string) []string {
	result := []string{}
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] != '#' {
			j++
		}

		l, _ := strconv.Atoi(s[i:j])
		fmt.Println("length", l)
		start := j + 1
		end := start + l
		result = append(result, s[start:end])
		i = end
	}
	return result
}
