package hashing

import "fmt"

func countingTriangles() {
	input := [][]int{{1, 2}, {2, 1}, {2, 2}, {2, 3}, {3, 2}}
	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for _, val := range input {
		xValue := val[0]
		yValue := val[1]

		xMap[xValue] = xMap[xValue] + 1
		yMap[yValue] = yMap[yValue] + 1
	}

	count := 0
	for _, val := range input {
		x := val[0]
		y := val[1]
		fx := xMap[x]
		fy := yMap[y]

		count += (fx - 1) * (fy - 1)
	}

	fmt.Println(count)
}
