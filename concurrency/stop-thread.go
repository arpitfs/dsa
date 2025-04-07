package concurrency

import (
	"fmt"
	"time"
)

func workerProcess(ch chan bool) {
	fmt.Println("working")

	for {
		select {
		case <-ch:
			fmt.Println("Completed Processing")
			return
		default:
			fmt.Println("Processing")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerMain() {
	ch := make(chan bool)
	go workerProcess(ch)

	time.Sleep(3 * time.Second)
	ch <- true

	time.Sleep(1 * time.Second)
	fmt.Println("Completed")
}
