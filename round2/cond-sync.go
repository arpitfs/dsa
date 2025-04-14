package round2

import (
	"fmt"
	"sync"
	"time"
)

var (
	data   int
	syncMu sync.Mutex
	cond   = sync.NewCond(&syncMu)
)

func conditionalSync() {
	go produce()
	go consume()

	time.Sleep(15 * time.Second)
}

func produce() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		syncMu.Lock()
		data = i
		fmt.Println("Produced: ", data)
		cond.Signal()
		syncMu.Unlock()
	}
}

func consume() {
	for {
		syncMu.Lock()
		cond.Wait()
		fmt.Println("Concumed: ", data)
		syncMu.Unlock()
	}
}
