package bitwise

import "fmt"

func bitOperations() {
	input := 5
	bit := 2

	getIBit(input, bit)
	setIBit(input, bit)
	clearIBit(input, bit)

}

func getIBit(input, bit int) {
	mask := 1 << bit
	if input&mask > 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

func setIBit(input, bit int) {
	mask := 1 << bit
	fmt.Println(input | mask)
}

func clearIBit(input, bit int) {
	mask := 1 << bit
	invertedMask := ^mask
	fmt.Println(input & invertedMask)
}
