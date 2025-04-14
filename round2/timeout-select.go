package round2

import (
	"fmt"
	"time"
)

func timeoutSelectMain() {
	timeout := time.After(4 * time.Second)
	ch := make(chan int, 5)
	go distribute(ch)

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("Timeout")
			return
		}
	}
}

func distribute(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * time.Duration(i))
		ch <- i
	}
	close(ch)
}
