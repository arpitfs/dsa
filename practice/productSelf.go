package practice

import "fmt"

func productWithoutSelf() {
	input := []int{1, 2, 3, 4}
	length := len(input)

	prefixArray := make([]int, length)
	prefixArray[0] = 1
	for i := 1; i < length; i++ {
		prefixArray[i] = prefixArray[i-1] * input[i-1]
	}

	fmt.Println(prefixArray)

	suffixProduct := 1
	// suffixArray := make([]int, length)
	// suffixArray[length-1] = 1

	for i := length - 1; i >= 0; i-- {
		prefixArray[i] = prefixArray[i] * suffixProduct
		suffixProduct = suffixProduct * input[i]
	}

	// for i := len(suffixArray) - 2; i >= 0; i-- {
	// 	suffixArray[i] = suffixArray[i+1] * input[i+1]
	// }

	// fmt.Println(suffixArray)

	// resultArray := make([]int, length)
	// for i := 0; i < length; i++ {
	// 	resultArray[i] = prefixArray[i] * suffixArray[i]
	// }
	fmt.Println(prefixArray)
	// [1,1,2,6]  // [1,4,12,24]
	// result :[1,4,24,]

}
