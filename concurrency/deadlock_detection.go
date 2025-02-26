package concurrency

import (
	"fmt"
	"sync"
)

var mu1 sync.Mutex
var mu2 sync.Mutex

func detectConcurrency() {
	var wg sync.WaitGroup
	wg.Add(2)
	go threadOne(&wg)
	go threadTwo(&wg)

	wg.Wait()

}

func threadOne(wg *sync.WaitGroup) {
	defer wg.Done()
	mu1.Lock()
	fmt.Println("Thread One in action")

	mu2.Lock()
	fmt.Println("Thread one on mu2")

	mu2.Unlock()
	mu1.Unlock()

}

func threadTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	mu1.Lock()
	fmt.Println("Thread Two in action")

	mu2.Lock()
	fmt.Println("Thread Two on mu2")

	mu2.Unlock()
	mu1.Unlock()

}
