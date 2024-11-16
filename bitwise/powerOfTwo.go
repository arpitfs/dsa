package bitwise

import "fmt"

func checkPowerOfTwo() {
	input := 15

	if input&(input-1) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
