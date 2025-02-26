package concurrency

import "fmt"

// Write a producer-consumer example where the producer closes the channel after sending values, and the consumer correctly handles it.

func producerConsumer() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
