package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Use sync.Cond to implement a producer-consumer scenario where the consumer waits for the producer to send data.

var (
	data   int
	syncMu sync.Mutex

	cond = sync.NewCond(&syncMu)
)

func condSyc() {

	go producerSync()
	go consumerSync()
	time.Sleep(12 * time.Second)

}

func producerSync() {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		syncMu.Lock()
		data = i
		fmt.Println("Produced: ", data)
		cond.Signal()
		syncMu.Unlock()
	}
}

func consumerSync() {
	for {
		syncMu.Lock()
		cond.Wait()
		fmt.Println("Consumed Data: ", data)
		syncMu.Unlock()
	}
}
