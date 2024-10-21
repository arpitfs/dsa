package stringMain

import "fmt"

func anagram() {
	input1 := "anagra"
	input2 := "ngaram"
	fmt.Println(validAnagram(input1, input2))

}

func validAnagram(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}

	tracker := make([]int, 26)
	for i, _ := range input1 {
		input1Value := rune(input1[i]) - rune('a')
		tracker[input1Value] = tracker[input1Value] + 1
		input2Value := rune(input2[i]) - rune('a')
		tracker[input2Value] = tracker[input2Value] - 1
	}

	for i, _ := range tracker {
		if tracker[i] != 0 {
			return false
		}
	}

	return true
}
