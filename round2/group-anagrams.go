package round2

import (
	"fmt"
	"sort"
)

func groupAnangram() {
	input := []string{"eat", "tea", "ate", "tan", "nat", "cat"}

	result := make(map[string][]string)

	for i, _ := range input {
		sortedString := sortInput(input[i])
		result[sortedString] = append(result[sortedString], input[i])
	}

	fmt.Println(result)
}

func sortInput(s string) string {
	convertedString := []rune(s)
	sort.Slice(convertedString, func(i, j int) bool {
		return convertedString[i] < convertedString[j]
	})
	return string(convertedString)
}
