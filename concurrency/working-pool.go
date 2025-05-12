package concurrency

import (
	"fmt"
	"sync"
)

func workingPool() {
	const nJobs = 20
	const threads = 3

	var wg sync.WaitGroup
	s := make(chan struct{}, threads)

	for j := 1; j <= nJobs; j++ {
		wg.Add(1)
		s <- struct{}{}

		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-s
		}(j)
	}

	wg.Wait()
}
