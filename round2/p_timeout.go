package round2

import (
	"fmt"
	"time"
)

func timeouts() {
	ch := make(chan int, 4)

	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			ch <- i
		}
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-time.After(2 * time.Second):
			return
		}
	}
}
