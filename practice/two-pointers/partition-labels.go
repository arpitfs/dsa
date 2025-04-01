package twopointers

import "fmt"

//Divide a string into partitions where each letter appears in at most one part.
// input = "ababcbacadefegdehijhklij"
// result = [9, 7, 8]
func PartitionLables() {
	input := "ababcbacadefegdehijhklij"
	tracker := make(map[rune]int)

	for i, v := range input {
		tracker[v] = i
	}

	var result []int
	start, end := 0, 0
	for i, ch := range input {
		if tracker[ch] > end {
			end = tracker[ch]
		}

		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	fmt.Println(result)

}
