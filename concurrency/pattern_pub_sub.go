package concurrency

import (
	"fmt"
	"time"
)

func patternPubSub() {
	input := make(chan string)
	sub1 := make(chan string)
	sub2 := make(chan string)

	go broadcaster(input, sub1, sub2)

	go func() {
		for msg := range sub1 {
			fmt.Println("Sub1 received:", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Println("Sub2 received:", msg)
		}
	}()

	input <- "Hello"
	input <- "World"
	close(input)

	time.Sleep(time.Second)
}

func broadcaster(input <-chan string, subscribers ...chan string) {
	for msg := range input {
		for _, sub := range subscribers {
			sub <- msg
		}
	}
	for _, sub := range subscribers {
		close(sub)
	}
}
