package main

import (
	"fmt"
	"sort"
)

type Result struct {
	firstValue  int
	secondValue int
	thirdValue  int
}

func threeSum() {
	input := []int{-1, 0, 1, 2, -1, -4}
	result := make(map[int]Result)
	sort.Ints(input) // -4,-1,-1,0,1,2
	for i := 0; i <= len(input)-3; i++ {
		if i == 0 || input[i] != input[i-1] {
			left := i + 1
			right := len(input) - 1
			target := 0 - input[i]
			for left < right {
				if input[left]+input[right] == target {
					r := Result{}
					r.firstValue = input[i]
					r.secondValue = input[left]
					r.thirdValue = input[right]
					result[r.firstValue] = r
					for left < len(input)-1 && input[left] == input[left+1] {
						left++
					}

					for right > 0 && input[right] == input[right-1] {
						right--
					}
					left++
					right--

				} else if input[left]+input[right] < target {
					left++
				} else {
					right--
				}

			}
		}

	}
	fmt.Println(result)

}
