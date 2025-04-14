package round2

import (
	"fmt"
	"sync"
)

func fanInOut() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go workerDo(i, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}

func workerDo(i int, jobs, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println(j)
		results <- j * 2
	}
}
