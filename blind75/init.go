package blind75

import "fmt"

const i = 1

func init() {
	fmt.Println("Init FIrst", i)
}

func init() {
	fmt.Println("Init Second")
}

func test() {
	const i = 0
	fmt.Println("Test", i)
}
