package concurrency

import (
	"fmt"
	"time"
)

//

func rateLimitor() {
	reqChannel := make(chan int, 10)
	go processor(reqChannel, time.Second)

	for i := 0; i < 5; i++ {
		reqChannel <- i
		fmt.Println("Request Send")
	}
	time.Sleep(6 * time.Second)
}

func processor(channel <-chan int, rate time.Duration) {
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	for req := range channel {
		<-ticker.C
		processRequest(req)
	}
}

func processRequest(data int) {
	fmt.Println(data)

}
