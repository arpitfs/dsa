package round2

import (
	"fmt"
	"sync"
)

func channelWorking() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendNumbers(ch, &wg)
	go printNumbers(ch, &wg)
	wg.Wait()
}

func sendNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func printNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		fmt.Println(c)
	}
}
