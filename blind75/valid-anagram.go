package blind75

import "fmt"

func validAnagram() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(checkAnagram(s, t))
}

func checkAnagram(s, t string) bool {
	tracker := make([]int, 26)
	for i, _ := range s {
		input1 := s[i] - 'a'
		tracker[input1]++
		input2 := t[i] - 'a'
		tracker[input2]--
	}

	for _, v := range tracker {
		if v != 0 {
			return false
		}
	}
	return true
}
