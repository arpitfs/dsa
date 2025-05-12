package concurrency

import (
	"fmt"
	"time"
)

func timeoutWait() {
	processCh := make(chan int, 20)
	go process(processCh)

	for i := 0; i < 10; i++ {
		//time.Sleep(2 * time.Second)
		processCh <- i
	}

	time.Sleep(5 * time.Second)

}

func process(p chan int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case msg := <-p:
			fmt.Println(msg)
		case <-time.After(2 * time.Second):
			fmt.Println("TImeout")
			return
		}

	}
}
