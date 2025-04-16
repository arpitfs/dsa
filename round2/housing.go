package round2

import "fmt"

type House struct {
	start int
	end   int
}

func housing() {
	input := []int{1, 3, 2, 1, 4, 1, 3, 2, 1, 1, 2}
	target := 8
	left, right := 0, 0
	currentSum := 0
	var house []House

	for right < len(input) {
		currentSum += input[right]

		for currentSum > target && left <= right {
			currentSum -= input[left]
			left++
		}

		if currentSum == target {
			house = append(house, House{start: left, end: right})
		}

		right++
	}

	fmt.Println(house)
}
