package concurrency

import (
	"fmt"
	"time"
)

//  Demonstrate the difference between buffered and unbuffered channels with an example.

func channels() {
	unbufferedChannel()
	bufferedChannel()
}

func unbufferedChannel() {
	ch := make(chan string)
	go func() {
		fmt.Println("Sending Value")
		ch <- "Unbufferd Channel"

	}()
	time.Sleep(time.Second)
	fmt.Println(<-ch)

}

func bufferedChannel() {
	ch := make(chan string, 2)
	fmt.Println("Sending Value")
	ch <- "bufferd"
	ch <- "channel"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
