package concurrency

import "fmt"

func processing() {
	numChan := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i < 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	go func() {
		s := 0
		for n := range numChan {
			s += n
		}
		fmt.Println(s)
		done <- true
	}()

	<-done
}
