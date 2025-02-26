package concurrency

import (
	"fmt"
	"sync"
)

//Write a program that spawns 10 goroutines, each printing its index, and ensure the main function waits for all to complete.

func printThreads() {
	var wg sync.WaitGroup
	threads := 10
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go printIndex(i, &wg)
	}

	wg.Wait()
}

func printIndex(index int, wg *sync.WaitGroup) {
	fmt.Println("Thread No:", index)
	wg.Done()
}
