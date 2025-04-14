package round2

import (
	"fmt"
	"time"
)

func rateLimitor() {
	ch := make(chan int, 10)
	go process(ch, time.Second)

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Request Send")
	}

	time.Sleep(10 * time.Second)
}

func process(ch chan int, rate time.Duration) {
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	for c := range ch {
		<-ticker.C
		fmt.Println(c)
	}
}
