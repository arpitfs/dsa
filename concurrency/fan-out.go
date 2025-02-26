package concurrency

import (
	"fmt"
	"math"
	"sync"
)

// Distribute tasks from one channel to multiple worker goroutines and collect results.
func fanOut() {

	ch := make(chan int, 5)
	var muLock sync.Mutex
	result := make(map[int]int)

	var wg sync.WaitGroup
	numWorkers := 7
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &result, &muLock, &wg)
	}
	distributor(ch)
	wg.Wait()
	//	time.Sleep(2 * time.Second)
	fmt.Println(result)

}

func distributor(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i + 1
	}
	close(ch)
}

func worker(data int, ch chan int, result *map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m := <-ch
	data1 := math.Pow(float64(m), 2)
	mu.Lock()
	(*result)[int(data1)] = int(data1)
	mu.Unlock()
}
