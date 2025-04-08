package concurrency

import (
	"fmt"
	"sync"
)

func FanInOut() {
	jobs := make(chan int, 5)
	result := make(chan int, 5)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerFan(i, jobs, result, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(result)
	}()

	for result := range result {
		fmt.Println("Result:", result)
	}

}

func workerFan(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
}
