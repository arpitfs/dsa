package concurrency

import (
	"fmt"
	"sync"
)

// Use sync.Once to ensure a function is only executed once, even when called by multiple goroutines.

var (
	muOnce       sync.Once
	initialValue int
)

func initialize() {
	initialValue = 2
	fmt.Println("Intializing", initialValue)
}

func OnceWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Thread Starting")
	muOnce.Do(initialize)

}

func syncOnce() {
	var wg sync.WaitGroup
	threads := 5
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go OnceWorker(&wg)
	}

	wg.Wait()
}
