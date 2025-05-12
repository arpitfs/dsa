package concurrency

import "fmt"

func pipelineCheck() {
	sq := make(chan int)
	doub := make(chan int)

	go func() {
		for i := 1; i < 5; i++ {
			sq <- i * i
		}
		close(sq)
	}()

	go func() {
		for s := range sq {
			doub <- 2 * s
		}
		close(doub)
	}()

	for r := range doub {
		fmt.Println(r)
	}
}
