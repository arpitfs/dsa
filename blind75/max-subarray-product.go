package blind75

import "fmt"

func maxSubarrayProduct() {
	input := []int{2, 3, -2, 4}
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

func minimum(c1, v1 int) int {
	if c1 < v1 {
		return c1
	}

	return v1
}
