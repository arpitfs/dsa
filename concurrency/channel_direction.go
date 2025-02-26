package concurrency

import "fmt"

// Implement a function that only sends data to a channel and another that only receives from it using channel direction syntax.
func channelDirection() {
	ch := make(chan int, 2)
	go send(ch)

	receriver(ch)

}

func send(ch chan<- int) {
	for i := 0; i < 2; i++ {
		ch <- i + 1
	}
	close(ch)
}

func receriver(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}
