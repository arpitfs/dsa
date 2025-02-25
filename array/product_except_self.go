package array

import "fmt"

func productExeceptSelf() {
	input := []int{1, 2, 3, 4}
	length := len(input)
	prefixSubArray := make([]int, length)
	prefixSubArray[0] = 1
	for i := 1; i < length; i++ {
		prefixSubArray[i] = prefixSubArray[i-1] * input[i-1]

	}

	suffixProduct := 1
	for i := length - 1; i >= 0; i-- {
		prefixSubArray[i] = prefixSubArray[i] * suffixProduct
		suffixProduct = suffixProduct * input[i]
	}

	fmt.Println(prefixSubArray)

}
