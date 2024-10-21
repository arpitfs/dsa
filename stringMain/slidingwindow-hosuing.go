package stringMain

import "fmt"

func houses() {
	input := []int{1, 3, 2, 1, 4, 1, 3, 2, 1, 1, 2}
	target := 8
	getHouses(input, target)
}

func getHouses(input []int, target int) {
	i := 0
	j := 0
	c := 0
	for j < len(input) {
		c += input[j]
		j++
		for c > target && i < j {
			c = c - input[i]
			i++
		}

		if c == target {
			fmt.Println(i, j-1)
		}
	}

}
