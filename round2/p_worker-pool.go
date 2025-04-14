package round2

import (
	"fmt"
	"sync"
)

func workerPoolService() {
	tasks := make(chan int, 10)
	maxThreads := 2
	var wg sync.WaitGroup
	go func() {
		for i := 0; i < 8; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		go working(i, tasks, &wg)
	}

	wg.Wait()
}

func working(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range ch {
		fmt.Println(t)
	}
}
