package round2

import (
	"fmt"
	"sync"
)

var (
	mOnce   sync.Once
	initial int
)

func initialize() {
	initial = 2
	fmt.Println("Value initialized")
}

func processWork(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Thread Starting")
	mOnce.Do(initialize)
}
func syncOnceMain() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processWork(&wg)
	}
	wg.Wait()
}
