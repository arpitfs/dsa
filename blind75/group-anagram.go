package blind75

import (
	"fmt"
	"sort"
)

func groupAnagram() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := make(map[string][]string)
	for i, _ := range input {
		sortedInput := sortString(input[i])
		result[sortedInput] = append(result[sortedInput], input[i])
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
