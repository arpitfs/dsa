package round2

import "fmt"

func maxSubArraySum() {
	input := []int{5, 4, -1, 7, 8}
	ans := input[0]
	currentSum := input[0]
	for i := 1; i < len(input); i++ {
		currentSum = maximum(input[i], currentSum+input[i])
		ans = maximum(ans, currentSum)
	}

	fmt.Println(ans)
}
