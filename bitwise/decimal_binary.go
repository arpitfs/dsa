package bitwise

import "fmt"

func convertNumber() {
	number := 13
	power := 0
	ans := 0
	for number > 0 {
		rem := number % 2
		number = number / 2
		ans = ans + intPow(10, power)*rem
		power++
	}
	fmt.Println(ans)
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
