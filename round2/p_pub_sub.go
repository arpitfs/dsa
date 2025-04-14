package round2

import (
	"fmt"
	"time"
)

func pubSub() {
	input := make(chan string)
	sub1 := make(chan string)
	sub2 := make(chan string)

	go broadcaste(input, sub1, sub2)
	go func() {
		for m := range sub1 {
			fmt.Println("Sub1", m)
		}
	}()

	go func() {
		for m := range sub2 {
			fmt.Println("Sub2s", m)
		}
	}()

	input <- "Hi"
	input <- "Hello"
	close(input)
	time.Sleep(time.Second * 3)
}

func broadcaste(input chan string, subsribers ...chan string) {
	for m := range input {
		for _, s := range subsribers {
			s <- m
		}
	}

	for _, s := range subsribers {
		close(s)
	}
}
