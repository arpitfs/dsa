package round2

import "fmt"

func validAnagramProblem() {
	input1 := "anagra"
	input2 := "ngaraa"
	fmt.Println(validAnagram(input1, input2))
}

func validAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	checker := make([]int, 26)

	for i, _ := range s1 {
		s1Value := s1[i] - 'a'
		checker[s1Value]++
		s2Value := s2[i] - 'a'
		checker[s2Value]--
	}

	for i, _ := range checker {
		if checker[i] != 0 {
			return false
		}

	}
	return true
}
