package concurrency

import "fmt"

// Merge multiple input channels into a single output channel.

func fanIn() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	result := make(chan int)

	data1 := []int{1, 2, 3}
	data2 := []int{4, 5, 6}
	data3 := []int{7, 8, 9}

	go sendData(ch1, data1)
	go sendData(ch2, data2)
	go sendData(ch3, data3)

	go receiver([]chan int{ch1, ch2, ch3}, result)

	for i := 0; i < 10; i++ {
		fmt.Println(<-result)
	}

}

func sendData(ch chan<- int, data []int) {
	for _, val := range data {
		ch <- val
	}
	close(ch)

}

func receiver(channels []chan int, out chan<- int) {

	for {
		select {
		case msg := <-channels[0]:
			out <- msg
		case msg := <-channels[1]:
			out <- msg
		case msg := <-channels[2]:
			out <- msg
		}
	}

}
