package concurrency

import (
	"fmt"
	"time"
)

// Write a function that waits for input from a channel but times out if it doesn't receive anything within 2 seconds.

func timeoutSelect() {
	timeout := time.After(4 * time.Second)
	ch := make(chan int, 5)
	go distributorWorkload(ch)
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("Exiting")
			return
		}
	}

}

func distributorWorkload(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * time.Duration(i))
		ch <- i
	}
	close(ch)
}
