package round2

import "fmt"

func closureWorking() {
	incrementerOne := closureFunction()
	fmt.Println(incrementerOne(1))
	fmt.Println(incrementerOne(1))
	incrementerSecond := closureFunction()
	fmt.Println(incrementerSecond(1))
	fmt.Println(incrementerSecond(1))

}

func closureFunction() func(int) int {
	count := 0
	return func(x int) int {
		count += x
		return count
	}
}
