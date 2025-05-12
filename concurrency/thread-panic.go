package concurrency

import (
	"fmt"
	"time"
)

func threadPanic() {
	go recoverSystem()

	time.Sleep(4 * time.Second)
}

func recoverSystem() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic")
		}
	}()
	fmt.Println("WOrking")
	time.Sleep(1 * time.Second)
	panic("Panic inside thread")
}
