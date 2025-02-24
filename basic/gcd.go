package basic

import "fmt"

func gcd() {

	input1 := 24
	input2 := 36

	fmt.Println(findGcd(input1, input2))

}

func findGcd(input1, input2 int) int {
	if input2 == 0 {
		return input1
	}

	return findGcd(input2, input1%input2)
}
