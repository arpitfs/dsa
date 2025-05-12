package concurrency

import "fmt"

func bufferedProcessing() {
	numChan := make(chan int)
	//done := make(chan bool)

	go func() {
		for i := 1; i < 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	s := 0
	for {
		v, ok := <-numChan
		if ok {
			s += v
		} else {
			fmt.Println("no values")
			break
		}
	}
	fmt.Println(s)

	//<-done
}
