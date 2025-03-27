package practice

import (
	"fmt"
	"sync"
)

func sumCon() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go sum(ch, &wg)
	go rec(ch, &wg)
	wg.Wait()

}

func sum(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func rec(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for v := range ch {
		sum += v
	}

	fmt.Println(sum)
}
