package concurrency

import (
	"fmt"
	"sync"
)

// Use sync.Mutex to safely increment a counter shared across multiple goroutines.

func counter() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	value := 1
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go increment(&value, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(value)

}

func increment(i *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*i = *i + 1
	mu.Unlock()
}
