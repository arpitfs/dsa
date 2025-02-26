package concurrency

import (
	"fmt"
)

// Implement a program that listens to two channels and prints messages from whichever channel receives data first.
func selectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producerChannelOne(ch1, "Channel One")
	go producerChannelTwo(ch2, "Channel Two")

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

}

func producerChannelOne(ch chan string, data string) {
	ch <- data
}

func producerChannelTwo(ch chan string, data string) {
	ch <- data

}
