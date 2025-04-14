package round2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64 = 0

func atomicCounter() {
	thread := 10
	var wg sync.WaitGroup
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
		//go increment(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&count, 1)
	fmt.Println(count)
}
