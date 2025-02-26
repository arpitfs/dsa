package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Use sync/atomic to safely increment a counter across multiple goroutine

var count int64

func atomicCounter() {
	thread := 10
	var wg sync.WaitGroup

	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
