package round2

import (
	"fmt"
	"sync"
)

func channelWorkingDirection() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendNumbers(ch, &wg)
	go printNumbers(ch, &wg)
	wg.Wait()
}

func sendNumbersDirection(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func printNumbersDirection(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		fmt.Println(c)
	}
}
