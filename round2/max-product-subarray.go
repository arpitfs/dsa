package round2

import "fmt"

func findMaxProductSubarray() {
	input := []int{-2, 3, -4}
	result := input[0]
	maxProduct := input[0]
	minProduct := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		minProduct = minimum(input[i], minProduct*input[i])
		maxProduct = maximum(input[i], maxProduct*input[i])

		result = maximum(result, maxProduct)
	}

	fmt.Println(result)
}

func minimum(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}
