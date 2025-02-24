package array

import "fmt"

func lcm() {
	input1 := 12
	input2 := 18
	fmt.Println(calculateLCM(input1, input2))

}

func calculateLCM(input1, input2 int) int {
	lcm := (input1 * input2) / findGcd(input1, input2)
	return lcm
}
