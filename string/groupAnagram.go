package string

import (
	"fmt"
	"sort"
)

func groupAnangram() {
	input := []string{"eat", "tea", "ate", "tan", "nat", "cat"}
	result := make(map[string][]string)

	for i, _ := range input {
		sortedInput := sortString(input[i])
		if _, e := result[sortedInput]; e {
			result[sortedInput] = append(result[sortedInput], input[i])
		} else {
			result[sortedInput] = append(result[sortedInput], input[i])
		}
	}

	fmt.Println(result)
}

func sortString(input string) string {
	convertedString := []rune(input)
	sort.Slice(convertedString, func(i, j int) bool {
		return convertedString[i] < convertedString[j]
	})

	return string(convertedString)
}
