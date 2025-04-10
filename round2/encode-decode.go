package round2

import (
	"fmt"
	"strconv"
	"strings"
)

func encodeDecode() {
	input := []string{"hello", "world", "go"}
	result := encode(input)
	decodedResult := decode(result)
	fmt.Println(decodedResult)
}

func decode(input string) []string {
	result := []string{}
	i := 0
	for i < len(input) {
		j := i
		for j < len(input) && input[j] != '#' {
			j++
		}
	l, _ := strconv.Atoi(input[i:j])
		start := j + 1
		end := start + l

		result = append(result, input[start:end])
		i = end
	}
	return result
}

func encode(s []string) string {
	var result strings.Builder

	for _, v := range s {
		result.WriteString(strconv.Itoa(len(v)))
		result.WriteString("#")
		result.WriteString(v)
	}

	return result.String()
}
