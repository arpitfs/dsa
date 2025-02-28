package array

import (
	"fmt"
	"math"
)

// Find the Largest Contiguous Subarray Product
// Input: [2,3,-2,4]

// Output: 6

func largestContinousProduct() {
	input := []int{2, 3, -2, 4}
	maxProd, minProd, result := input[0], input[0], input[0]

	for i := 1; i < len(input); i++ {
		current := input[i]
		if current < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = int(math.Max(float64(current), float64(current*maxProd)))
		minProd = int(math.Max(float64(current), float64(current*minProd)))
		result = int(math.Max(float64(result), float64(maxProd)))
	}
	fmt.Println(result)

}
