package concurrency

import (
	"fmt"
)

func closure() {
	firstClosure := add()
	fmt.Println(firstClosure(1))
	fmt.Println(firstClosure(1))
	fmt.Println(firstClosure(1))
	sClosure := add()
	fmt.Println(sClosure(1))
	fmt.Println(sClosure(1))
	fmt.Println(sClosure(1))

}

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
