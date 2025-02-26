package concurrency

import (
	"fmt"
	"sync"
)

// Implement a scenario where multiple readers can read a shared resource concurrently but only one writer can modify it at a time using sync.RWMutex.

var sharedResource int
var mu sync.RWMutex

func readWriteMutex() {
	var wg sync.WaitGroup
	wg.Add(6)
	go writer(&wg)
	go writer(&wg)
	go read(&wg)
	go read(&wg)
	go writer(&wg)
	go read(&wg)
	wg.Wait()

}

func writer(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	fmt.Println("Updating resource")
	sharedResource = sharedResource + 1
	mu.Unlock()

}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	fmt.Println("Reading resource:", sharedResource)
	mu.Unlock()
}
