package round2

import (
	"fmt"
	"time"
)

func stopThread() {
	ch := make(chan bool)
	go startProcessing(ch)
	time.Sleep(3 * time.Second)
	ch <- true
	time.Sleep(1 * time.Second)
	fmt.Println("Completed")
}

func startProcessing(ch chan bool) {
	fmt.Println("Work started")

	for {
		select {
		case <-ch:
			fmt.Println("Processing Completed")
		default:
			fmt.Println("Processing")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
