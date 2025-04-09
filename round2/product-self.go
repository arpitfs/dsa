package round2

import "fmt"

func productSelfQuestion() {
	input := []int{1, 2, 3, 4}

	prefixArray := make([]int, len(input))
	prefixArray[0] = 1
	for i := 1; i < len(input); i++ {
		prefixArray[i] = prefixArray[i-1] * input[i-1]
	}

	suffixProduct := 1

	for i := len(input) - 1; i >= 0; i-- {
		prefixArray[i] = prefixArray[i] * suffixProduct
		suffixProduct = suffixProduct * input[i]
	}

	fmt.Println(prefixArray)

}
