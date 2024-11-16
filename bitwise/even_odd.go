package bitwise

import "fmt"

func checkEvenOdd() {
	input := 2
	if input&1 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
