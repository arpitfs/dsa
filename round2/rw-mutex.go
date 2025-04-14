package round2

import (
	"fmt"
	"sync"
)

var (
	resource int
	mu       sync.RWMutex
)

func readWriteMutex() {
	var wg sync.WaitGroup
	wg.Add(5)

	go writer(&wg)
	go writer(&wg)
	go reader(&wg)
	go writer(&wg)
	go reader(&wg)
	wg.Wait()
}

func writer(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	fmt.Println("updating resource")
	resource += 1
	mu.Unlock()
}

func reader(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	fmt.Println("Reading resource ", resource)
	mu.Unlock()
}
