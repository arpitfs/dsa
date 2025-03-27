package practice

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func channelConcurrency() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go inc(&wg, &mu)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go dec(&wg, &mu)
	}

	wg.Wait()
	fmt.Println(counter)
}

func inc(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	counter += 1
	mu.Unlock()
}

func dec(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	counter -= 1
	mu.Unlock()
}
