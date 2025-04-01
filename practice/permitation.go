package practice

// Check if a permutation of s1 exists in s2
import "fmt"

func CheckPermutationExists() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkData(s1, s2))
}

func checkData(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}

	s1Count := make([]int, 26)
	s2Count := make([]int, 26)

	for i := 0; i < n; i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	if mathces(s1Count, s2Count) {
		return true
	}

	for i := n; i < m; i++ {
		s2Count[s2[i]-'a']++
		s2Count[s2[i-n]-'a']--

		if mathces(s1Count, s2Count) {
			return true
		}
	}
	return false
}

func mathces(s1Count, s2Count []int) bool {
	for i := 0; i < 26; i++ {
		if s1Count[i] != s2Count[i] {
			return false
		}
	}
	return true
}
