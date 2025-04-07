package concurrency

import "fmt"

func testPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover")
		}
	}()

	panic("Panic created")
	fmt.Println("NOt executed")
}
