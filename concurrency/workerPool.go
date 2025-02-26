package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Implement a worker pool where N workers process M tasks concurrently.

func pool() {
	task := make(chan int, 10)
	maxthreads := 2
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 8; i++ {
			task <- i
		}
		close(task)
	}()

	for i := 0; i < maxthreads; i++ {
		wg.Add(1)
		go workerPool(i, task, &wg)
	}

	wg.Wait()

}

func workerPool(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range ch {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(2 * time.Second)
	}
}
