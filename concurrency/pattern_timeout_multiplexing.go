package concurrency

import (
	"fmt"
	"time"
)

func Timeouts() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Data from channel"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
